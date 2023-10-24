# gb

Work in progress. Plan is to make a note taking cli program with `AES`.

Disclaimer I'm not any expert in any of this. I'm making for my own use.

# TO-DO

- [x] `init` command.
- [x] `new` command.
- [ ] `db` json.
- [ ] `open` command. // needs more work
- [ ] `edit` command.
- [ ] find more things

# Eample

```bash

~/fo/gb main
❯ go run . init # initializing
Bear in mind that you will never be able to change the pass :) so give a strong one
Enter the password:

defaltnotebook: default
editor: nvim

was written to "tmp/config.yml"

~/fo/gb main
❯ tree -a tmp/ # initial file tree
tmp/
├── config.yml
├── default
├── .gitignore
└── .key

2 directories, 3 files

~/fo/gb main
❯ go run . new # new note time stamp will be used

~/fo/gb main
❯ go run . new fo # new note with the name fo

~/fo/gb main
❯ tree -a tmp/
tmp/
├── config.yml
├── default
│   ├── 24-10-23_09-13-17-PM.md.enc
│   └── fo.md.enc
├── .gitignore
└── .key

2 directories, 5 files

~/fo/gb main
❯ cat tmp/default/24-10-23_09-13-17-PM.md.enc # encrypted hehe you cant read!
Kvv왾z][uùW&(Ϣ?n_I4l1Y⏎


```

