# Pointers and References in GoLang

- `*Foo` in a type definition is a pointer to `Foo`
- `*foo` dereferences the variable `foo` to provide access to the value of `foo`
- `&foo` returns the address of `foo`
- `&Foo{}` allocates memory for an object and returns the pointer.
- `new(Foo)` is same as `&Foo{}` _(the latter did not exist in early GoLang)_ 
- There is no `&Foo` 
- If `f1:=&Foo{}` and `f2:=Foo{}` then `*f1==f2` and `f1==&f2`
- `*f1==f2` compares content of structs _(mostly they are not, this would be `false`)_
- `*f1==f2` compares the pointers are the same



