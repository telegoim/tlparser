package fileid

import (
	"github.com/go-faster/errors"

	"github.com/gotd/td/bin"
)

// PhotoSizeSource represents photo metadata stored in file_id.
type PhotoSizeSource struct {
	Type      PhotoSizeSourceType
	VolumeID  int64
	LocalID   int
	Secret    int64
	PhotoSize string

	FileType      uint32
	ThumbnailType int32

	DialogID         int64
	DialogAccessHash int64

	StickerSetID         int64
	StickerSetAccessHash int64
	StickerVersion       int32
}

func (p *PhotoSizeSource) readLocalIDVolumeID(b *bin.Buffer) error {
	{
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read volume_id")
		}
		p.VolumeID = v
	}
	{
		v, err := b.Int()
		if err != nil {
			return errors.Wrap(err, "read local_id")
		}
		p.LocalID = v
	}
	return nil
}

func (p *PhotoSizeSource) readDialog(b *bin.Buffer) error {
	{
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read dialog_id")
		}
		p.DialogID = v
	}
	{
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read dialog_access_hash")
		}
		p.DialogAccessHash = v
	}
	return nil
}

func (p *PhotoSizeSource) readStickerSet(b *bin.Buffer) error {
	{
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read sticker_set_id")
		}
		p.StickerSetID = v
	}
	{
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read sticker_set_access_hash")
		}
		p.StickerSetAccessHash = v
	}
	return nil
}

func (p *PhotoSizeSource) decode(b *bin.Buffer, subVersion byte) error {
	if subVersion < 32 {
		{
			v, err := b.Long()
			if err != nil {
				return errors.Wrap(err, "read volume_id")
			}
			p.VolumeID = v
		}

		if subVersion < 22 {
			{
				v, err := b.Long()
				if err != nil {
					return errors.Wrap(err, "read secret")
				}
				p.Secret = v
			}
			{
				v, err := b.Int()
				if err != nil {
					return errors.Wrap(err, "read local_id")
				}
				p.LocalID = v
			}

			return nil
		}
	}

	var photoSizeType PhotoSizeSourceType
	if subVersion >= 4 {
		v, err := b.Int()
		if err != nil {
			return errors.Wrap(err, "read photo_size_type")
		}
		photoSizeType = PhotoSizeSourceType(v)
	}
	if photoSizeType < 0 || photoSizeType >= lastPhotoSizeSourceType {
		return errors.Errorf("unknown photo size source type %d", photoSizeType)
	}
	p.Type = photoSizeType

	switch photoSizeType {
	case PhotoSizeSourceLegacy:
		v, err := b.Long()
		if err != nil {
			return errors.Wrap(err, "read secret")
		}
		p.Secret = v
	case PhotoSizeSourceThumbnail:
		{
			v, err := b.Uint32()
			if err != nil {
				return errors.Wrap(err, "read file_type")
			}
			p.FileType = v
		}
		{
			v, err := b.Int32()
			if err != nil {
				return errors.Wrap(err, "read thumbnail_type")
			}
			p.ThumbnailType = v
		}
	case PhotoSizeSourceDialogPhotoBig, PhotoSizeSourceDialogPhotoSmall:
		if err := p.readDialog(b); err != nil {
			return errors.Wrap(err, "read dialog")
		}
	case PhotoSizeSourceStickerSetThumbnail:
		if err := p.readStickerSet(b); err != nil {
			return errors.Wrap(err, "read sticker_set")
		}
	case PhotoSizeSourceFullLegacy:
		{
			v, err := b.Long()
			if err != nil {
				return errors.Wrap(err, "read volume_id")
			}
			p.VolumeID = v
		}
		{
			v, err := b.Long()
			if err != nil {
				return errors.Wrap(err, "read secret")
			}
			p.Secret = v
		}
		{
			v, err := b.Int()
			if err != nil {
				return errors.Wrap(err, "read local_id")
			}
			p.LocalID = v
		}
	case PhotoSizeSourceDialogPhotoBigLegacy, PhotoSizeSourceDialogPhotoSmallLegacy:
		if err := p.readDialog(b); err != nil {
			return errors.Wrap(err, "read dialog")
		}
		if err := p.readLocalIDVolumeID(b); err != nil {
			return errors.Wrap(err, "read legacy photo")
		}

	case PhotoSizeSourceStickerSetThumbnailLegacy:
		if err := p.readStickerSet(b); err != nil {
			return errors.Wrap(err, "read sticker_set")
		}
		if err := p.readLocalIDVolumeID(b); err != nil {
			return errors.Wrap(err, "read legacy photo")
		}

	case PhotoSizeSourceStickerSetThumbnailVersion:
		if err := p.readStickerSet(b); err != nil {
			return errors.Wrap(err, "read sticker_set")
		}
		{
			v, err := b.Int32()
			if err != nil {
				return errors.Wrap(err, "read sticker_version")
			}
			p.StickerVersion = v
		}
	}

	if subVersion < 32 && subVersion >= 22 {
		v, err := b.Int()
		if err != nil {
			return errors.Wrap(err, "read local_id")
		}
		p.LocalID = v
	}
	return nil
}
