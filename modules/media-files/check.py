import os
import csv
find_command = 'find /home -type f {} 2>/dev/null'

with open("media.csv", 'r') as media:
    reader = csv.reader(media)
    command = ""
    for row in reader:
        for i in row:
            command += f'-name "*{i.strip()}" -o '
    command = command[:-4]
    print(find_command.format(command))

# DELETE ALL MEDIA FILES COMMAND
# DO NOT TOUCH:
# find /home -type f -name "*.mp3" -o -name "*.wav" -o -name "*.aac" -o -name "*.flac" -o -name "*.ogg" -o -name "*.wma" -o -name "*.m4a" -o -name "*.aiff" -o -name "*.amr" -o -name "*.alac" -o -name "*.opus" -o -name "*.mid" -o -name "*.midi" -o -name "*.ra" -o -name "*.aif" -exec rm -f {} +
