for i in `ls note | sort -n`; do
    echo "[`head -n 1 note/$i | sed 's/^# //'`](note/`echo $i | sed 's/.md$//'`)";
done > README.md