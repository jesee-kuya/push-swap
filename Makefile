CC = go
BINS = push-swap checker

all: $(BINS)

push-swap:
	$(CC) build -o push-swap

checker:
	$(CC) build -o checker

clean:
	rm -f $(BINS)

fclean: clean

re: fclean all
