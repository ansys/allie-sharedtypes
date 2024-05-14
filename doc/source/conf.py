"""Sphinx documentation configuration file."""

from datetime import datetime
import os

from ansys_sphinx_theme import ansys_logo_black, ansys_favicon, get_version_match

# Project information
project = "allie-sharedtypes"
copyright = f"(c) {datetime.now().year} ANSYS, Inc. All rights reserved"
author = "ANSYS, Inc."

# Read version from VERSION file
with open(os.path.join("..", "..", "VERSION"), "r") as f:
    version_file = f.readline().strip()

release = version = version_file
switcher_version = get_version_match(version_file)
cname = os.getenv("DOCUMENTATION_CNAME", "noname.com")
"""The canonical name of the webpage hosting the documentation."""

# Select desired logo, theme, and declare the html title
html_theme = "ansys_sphinx_theme"
html_short_title = html_title = project
html_logo = ansys_logo_black
html_favicon = ansys_favicon
html_context = {
    "github_user": "ansys",
    "github_repo": "allie-sharedtypes",
    "github_version": "main",
    "doc_path": "doc/source",
}
html_theme_options = {
    "github_url": "https://github.com/ansys/allie-sharedtypes",
    "additional_breadcrumbs": [
        ("Allie", "https://allie.docs.pyansys.com/"),
    ],
    "switcher": {
        "json_url": f"https://{cname}/versions.json",
        "version_match": switcher_version,
    },
    "check_switcher": False,
    "show_prev_next": False,
    "show_breadcrumbs": True,
    "use_edit_page_button": True,
}

# Sphinx extensions
extensions = [
    "sphinx_design",
    "sphinx_copybutton",
]


# Add any paths that contain templates here, relative to this directory.
templates_path = ["_templates"]

# The suffix(es) of source filenames.
source_suffix = ".rst"

# The master toctree document.
master_doc = "index"

source_suffix = {
    ".rst": "restructuredtext",
}

# The master toctree document.
master_doc = "index"
