# goorg
Golang lib to parse org-mode files

An attempt to parse org-mode into a struct. This only parses a single layer of headlines at a time, the idea being that you can run through the file to find the headline you want fairly quickly. This will eventually be useful for inserting a headline at arbitrary points in the file.
