import os
import shutil
import re

# Set environment variables for file paths
DOC_BUILD_HTML = "documentation-html"
SOURCE_FILE = os.path.join(DOC_BUILD_HTML, "api_reference", "test", "index.html")
SOURCE_DIRECTORY = "dist/pkg/github.com/ansys/allie-flowkit/pkg/"
REPLACEMENT_DIRECTORY = os.path.join(DOC_BUILD_HTML, "api_reference", "pkg")
ACTUAL_DIR = os.path.join(DOC_BUILD_HTML, "api_reference/")

# Check if REPLACEMENT_DIRECTORY exists, if not, create it
if not os.path.exists(REPLACEMENT_DIRECTORY):
    os.makedirs(REPLACEMENT_DIRECTORY)

# Remove existing content in the replacement directory
for filename in os.listdir(REPLACEMENT_DIRECTORY):
    file_path = os.path.join(REPLACEMENT_DIRECTORY, filename)
    try:
        if os.path.isfile(file_path) or os.path.islink(file_path):
            os.unlink(file_path)
        elif os.path.isdir(file_path):
            shutil.rmtree(file_path)
    except Exception as e:
        print(f'Failed to delete {file_path}. Reason: {e}')

# Move the source_directory content to the replacement_directory
for filename in os.listdir(SOURCE_DIRECTORY):
    shutil.move(os.path.join(SOURCE_DIRECTORY, filename), REPLACEMENT_DIRECTORY)

# Remove the index.html file in the replacement_directory
index_file_path = os.path.join(REPLACEMENT_DIRECTORY, "index.html")
if os.path.exists(index_file_path):
    os.remove(index_file_path)

# Process each HTML file in the replacement directory
for root, dirs, files in os.walk(REPLACEMENT_DIRECTORY):
    for file in files:
        if file.endswith(".html"):
            replacement_file = os.path.join(root, file)
            with open(replacement_file, 'r') as f:
                content = f.read()

            # Extract body content and perform replacements
            match = re.search(r'<body>([\s\S]*?)<\/body>', content)
            if match:
                replacementBodyContent = match.group(1)
                replacementBodyContent = re.sub(
                    r'<div class="top-heading" id="heading-wide"><a href="\/pkg\/github.com\/ansys\/allie-flowkit\/">GoPages \| Auto-generated docs<\/a><\/div>[\s\S]*?<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;<\/span><\/a>',
                    '',
                    replacementBodyContent
                )

                # Read source file and replace content
                with open(SOURCE_FILE, 'r') as f:
                    source_content = f.read()

                new_content = re.sub(
                    r'(<article class="bd-article" role="main">)[\s\S]*?(<\/article>)',
                    r'\1' + replacementBodyContent + r'\2',
                    source_content
                )

                # Write the modified content back to the replacement file
                with open(replacement_file, 'w') as f:
                    f.write(new_content)
            else:
                print(f'No <body> content found in {replacement_file}')

# Move the modified files back to the actual directory
for filename in os.listdir(REPLACEMENT_DIRECTORY):
    shutil.move(os.path.join(REPLACEMENT_DIRECTORY, filename), ACTUAL_DIR)
