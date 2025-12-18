---
description: Create application icons (macOS, Windows, Linux) from a source image using webp2icons
---

// turbo-all

1. Locate the source image (1024x1024 WebP or PNG).
   - If not provided, ask the user or search the current directory for high-resolution images.

2. Ensure `webp2icons` is available and built.
   - If the binary is missing, run `go build -o webp2icons .` in the tool directory.

3. Run the conversion command.
   - Example: `./webp2icons -i <input_path> -d <output_dir> -o <base_name>`
   - Use flags as directed by the user (e.g., `-mac`, `-win`, `-linux`).

4. Verify the generated files in the output directory.
   - `ls -R <output_dir>`

5. Report the result and the paths to the generated icons.
