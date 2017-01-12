package blobstore

type s3BucketIterator struct {
	blobCh    chan *Blob
	doneCh    chan struct{}
	blobCount uint
}

func (i *s3BucketIterator) Next() (*Blob, error) {
	if i.blobCh == nil {
		return nil, ErrIteratorDone
	}

	blob, ok := <-i.blobCh
	if !ok {
		i.blobCh = nil
		return nil, ErrIteratorDone
	}

	return blob, nil
}

func (i *s3BucketIterator) Done() {
	i.blobCh = nil
	close(i.doneCh)
}

func (i *s3BucketIterator) Count() uint {
	return i.blobCount
}
