package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ./repoInterface.StorageInterface -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i ./repoInterface.CacheInterface -o ./mocks/ -s "_minimock.go"
