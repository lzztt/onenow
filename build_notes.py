from glob import glob
from json import dump

notes = []

files = glob('note/*.md')
files.sort(key=lambda i: int(i[5:].split('_', 1)[0]))

for i in files:
    with open(i) as f:
        notes.append({'url': i[:-3], 'data': f.read()})

with open('src/gen/notes.json', 'w') as f:
    dump(notes, f, ensure_ascii=False)
