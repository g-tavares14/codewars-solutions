# Codewars Solutions

[![Perfil no Codewars](https://www.codewars.com/users/g-tavares14/badges/large)](https://www.codewars.com/users/g-tavares14)

Minhas soluções das katas do [Codewars](https://www.codewars.com), em **Go**. Uso o repositório para registrar minha evolução enquanto pratico a linguagem.

As soluções chegam aqui sozinhas: a cada **Submit** no Codewars, a extensão [codewars-sync](https://github.com/g-tavares14/codewars-sync) faz o commit do código e do enunciado da kata.

## Estrutura

```
8-kyu/
  opposite-number/
    README.md            ← enunciado da kata, com link para o Codewars
    opposite-number.go   ← minha solução
7-kyu/
  ...
```

As pastas seguem o rank de dificuldade do Codewars, que vai de **8 kyu** (iniciante) até **1 kyu** (mais difícil). Cada kata tem a própria pasta, com o enunciado e a solução.

## Formatação

O editor do Codewars não formata o código. A cada push com arquivos `.go`, o workflow [`gofmt`](.github/workflows/gofmt.yml) roda o `gofmt` e faz um commit `style: gofmt` quando algo muda, para as soluções seguirem sempre o padrão do Go.

## Rodando uma solução

As soluções usam `package kata`, como o Codewars exige, e não têm `main`. Para testar localmente, crie um `_test.go` na pasta da kata e rode:

```bash
go test ./8-kyu/opposite-number/
```

## Progresso

O badge acima mostra meu rank e honra atuais. A lista completa de katas resolvidas está no [meu perfil](https://www.codewars.com/users/g-tavares14/completed).
