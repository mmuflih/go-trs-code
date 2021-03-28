# Golang Transaction Code Generator
Transaction code generator with prefix and increment suffix

Install:

    go get -u github.com/mmuflih/go-trs-code

Usage:

    gocode.IncCode(prefix, code, force)

- pref: Prefix
- code: let it empty string if generate new code
- force: use true if you want to using old code and using false when you want to add suffix


Example:

    Normal Code => RG.210328.A
    ("RG", "", true)              => RG.210328.A
    ("RG", "RG.210328.A" , true)  => RG.210328.A
    ("RG", "", false)             => RG.210328.A
    ("RG", "RG.210328.A" , false) => RG.210328.B
    ("" , "", false)              => .210328.A
    ("" , "RG.210328.A" , false)  => RG.210328.B