# Define the target file (without extension) as an argument.
# For example, to process "example.md", run: make FILE=example
FILE ?=

# Define the output directory.
OUT_DIR := out

# Define the source directory.
SRC_DIR := images

# Pandoc options
PANDOC_OPTIONS := --template=template/eisvogel.tex --listings -V titlepage:true -V titlepage-background:$(SRC_DIR)/background9.pdf --resource-path=. --self-contained

# Default target
all: check-file
	@echo "Processing $(FILE)"
	@mkdir -p $(OUT_DIR)
	@cp -r $(SRC_DIR) $(OUT_DIR)
	@pandoc $(FILE).md -o $(OUT_DIR)/$(FILE).tex $(PANDOC_OPTIONS)

# Check if FILE variable is empty
check-file:
ifndef FILE
	$(error Please enter a filename without its extension using 'make FILE=<filename>')
endif

clean:
	@echo "Cleaning up"
	@rm -rf $(OUT_DIR)

.PHONY: all check-file clean
