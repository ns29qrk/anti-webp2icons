---
description: Create application icons (macOS, Windows, Linux) from a source image using webp2icons
---

// turbo-all

1. Locate the source image (1024x1024 WebP or PNG).
   - If not provided, ask the user or search the current directory for high-resolution images.

2. Ensure `webp2icons` is available and executable.
   - First, check if `webp2icons` is in the system PATH by running `which webp2icons`.
   - If not in PATH, check for the binary `./webp2icons` in the current directory.
   - If the binary is missing but `main.go` exists, run `go build -o webp2icons .` to build it.
   - If still missing, inform the user that the tool needs to be installed or built.

3. Run the conversion command using the discovered path to the binary.
   - Command: `webp2icons -i <input_path> -d <output_dir> -o <base_name>` (use `./webp2icons` if local binary).
   - Use flags as directed by the user (e.g., `-mac`, `-win`, `-linux`).

4. Verify the generated files in the output directory.
   - `ls -R <output_dir>`

5. Report the result and the paths to the generated icons.
