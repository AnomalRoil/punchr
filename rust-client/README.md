In order to build and start the rust-client, you must provide it with your API_KEY through the env variable:
```
cargo build
API_KEY="your api key" ./target/debug/rust-client
```

You can get logs by setting your env variable `RUST_LOG=debug` or `RUST_LOG=info`:
```
RUST_LOG=info API_KEY="your api key" ./target/debug/rust-client
```
