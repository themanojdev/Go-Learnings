#!/usr/bin/env bash

set -e

README="README.md"
TEMP_README=$(mktemp)
TEMP_CONCEPTS=$(mktemp)
TEMP_PROGRESS=$(mktemp)

echo "Starting README update..."

# Helper function: Extract clean title from markdown file
extract_title() {
    local file="$1"
    local title=""
    
    # Try to get title from file
    title=$(grep -m1 '^#' "$file" 2>/dev/null | sed 's/^#* *//')
    
    # Remove emojis and special characters, keep only text
    title=$(echo "$title" | sed 's/[^[:alnum:] \-\(\)]//g' | xargs)
    
    # Fallback to filename if no title found
    if [ -z "$title" ]; then
        title=$(basename "$file" .md)
    fi
    
    echo "$title"
}

# Generate Concepts Completed section
generate_concepts() {
    if [ ! -d "concepts" ]; then
        echo "Warning: concepts directory not found" >&2
        return
    fi
    
    find concepts -maxdepth 1 -type f -name "*.md" | sort | while read -r file; do
        if [ -f "$file" ]; then
            local name=$(basename "$file")
            local title=$(extract_title "$file")
            echo "- [x] **$name** - $title"
        fi
    done
}

# Generate Daily Progress section with deduplication
generate_progress() {
    if [ ! -d "daily-progress" ]; then
        echo "Warning: daily-progress directory not found" >&2
        return
    fi
    
    find daily-progress -mindepth 1 -maxdepth 1 | sort | while read -r item; do
        if [ -e "$item" ]; then
            local base=$(basename "$item")
            if [ -d "$item" ]; then
                echo "- [x] **${base}/**"
            else
                echo "- [x] $base"
            fi
        fi
    done | sort -u
}

# Get current date
TODAY=$(date +"%-d %B %Y")

# Generate concepts and progress to temp files
generate_concepts > "$TEMP_CONCEPTS"
generate_progress > "$TEMP_PROGRESS"

CONCEPTS_COUNT=$(wc -l < "$TEMP_CONCEPTS")
PROGRESS_COUNT=$(wc -l < "$TEMP_PROGRESS")

echo "Generated concepts: $CONCEPTS_COUNT lines"
echo "Generated progress: $PROGRESS_COUNT lines"

# Update README using awk with temp file approach
awk \
-v concepts_file="$TEMP_CONCEPTS" \
-v progress_file="$TEMP_PROGRESS" \
-v today="$TODAY" \
'
BEGIN {
    in_concepts = 0
    in_progress = 0
}
{
    # Handle Concepts Completed section header
    if ($0 ~ /^### Concepts Completed/) {
        print $0
        while ((getline line < concepts_file) > 0) {
            print line
        }
        close(concepts_file)
        in_concepts = 1
        in_progress = 0
        next
    }
    
    # Handle Daily Progress section header
    if ($0 ~ /^### Daily Progress/) {
        print $0
        while ((getline line < progress_file) > 0) {
            print line
        }
        close(progress_file)
        in_progress = 1
        in_concepts = 0
        next
    }
    
    # Exit current section when hitting next section header
    if ((in_concepts || in_progress) && $0 ~ /^###/) {
        in_concepts = 0
        in_progress = 0
    }
    
    # Skip old list items in active sections
    if ((in_concepts || in_progress) && $0 ~ /^- \[/) {
        next
    }
    
    # Update Last Updated timestamp
    if ($0 ~ /^Last Updated:/) {
        print "Last Updated: " today
        next
    }
    
    print $0
}
' "$README" > "$TEMP_README"

# Validate and replace original file
if [ -s "$TEMP_README" ]; then
    mv "$TEMP_README" "$README"
    echo "✓ README updated successfully on $TODAY"
else
    echo "✗ Error: Generated README is empty"
    rm -f "$TEMP_README"
    exit 1
fi

# Cleanup
rm -f "$TEMP_CONCEPTS" "$TEMP_PROGRESS"

exit 0