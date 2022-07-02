# l2-go

Layer 2 API

## Installation

run command from project root
```bash
go get github.com/trevatk/l2-go
```

## Usage

below is a snippet for creating a new block on the blockchain
```golang
	ctx := context.TODO()

	key := os.Getenv("APPLICATION_KEY")
	secret := os.Getenv("ORGANIZATION_SECRET")

	creds := credential.NewFromConfig(&credential.Config{
		Key:    key,
		Secret: secret,
	})

	l2 := l2.NewL2(creds)

	data := []byte("hello world")
	m := &model.Mock{
		Stake:      0.1,
		Difficulty: 1,
		Data:       data,
	}

	block, err := l2.MakeBlock(ctx, m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block hash: ", block.Hash)
```