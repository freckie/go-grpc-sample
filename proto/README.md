# /proto

- Directory for protobuf definitions.
- Separate each service into its own folder and place it under this folder like `/proto/tdoo`.
- All proto files in this folder are written by developers.

- Only Messages satisfying the following conditions are treated as Entities:
  - `*_ent.proto` file 
  - Message with the same name as the filename
  - (e.g., Message `User` in `user_ent.go`)
