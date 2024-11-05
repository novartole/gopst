build-deps:
	cargo build -r --manifest-path=launch-rs/Cargo.toml
	cp launch-rs/target/release/liblaunch_typst.dylib gopst/

run:
	go run -C gopst main.go
