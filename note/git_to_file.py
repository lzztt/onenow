from glob import glob
from subprocess import run


def main():
    for file in glob('*_*-*.md'):
        id, uuid = file[:-3].split('_', 2)

        commits = get_commits(file)

        saved = set()
        for time, hash in sorted(commits):
            git_path = get_hash_path(hash, f'^{id}_')
            data = get_content(git_path)

            if data not in saved:
                out = f'{uuid}_{time}.md'

                print(f"save {git_path} to {out}")
                with open(out, 'wb') as f:
                    f.write(data)

                saved.add(data)
            else:
                print(f"skip {git_path}")


def get_commits(file: str):
    res = run(
            f'git log --follow --format=%at_%h --date=unix {file}'.split(),
            capture_output=True,
            check=True,
        )

    commits = []
    for line in res.stdout.decode().strip().split("\n"):
        time, hash = line.split('_')
        commits.append((int(time), hash))

    return commits



def get_hash_path(hash: str, prefix: str) -> str:
    git = run(
        f'git show {hash}:note/'.split(),
        capture_output=True,
        check=True,
    )
    res = run(
        ['grep', prefix],
        input=git.stdout,
        capture_output=True,
        check=True,
        )

    return hash + ':note/' + res.stdout.decode().strip()


def get_content(hash_path: str) -> bytes:
    res = run(
        f'git show {hash_path}'.split(),
        capture_output=True,
        check=True,
    )

    return res.stdout


main()
