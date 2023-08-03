# knight tour decoder

https://github.com/Uchijo/knight-tour-encoder で生成したSAT問題を解いた結果を検証するためのツールです。
ナイトがどのようにグリッド上を動いたのかを見やすくすることができます。

## 動かし方

```sh
go run main.go <グリッドの一辺の長さ> <出力されたファイルのパス>
```

## 読み方

最初の `grid:` では、グリッドがどのような配置になっているかを確認できます。

続く `history:` では、ナイトがどのように動いたのかを確認できます。
timeが何手目であるか、coordがその時のグリッドの番号を示します。

```
grid:
  1   2   3   4   5 
  6   7   8   9  10 
 11  12  13  14  15 
 16  17  18  19  20 
 21  22  23  24  25 

history:
{rawValue:25 time:0 coord:25}
{rawValue:39 time:1 coord:14}
{rawValue:55 time:2 coord:5}
{rawValue:83 time:3 coord:8}
{rawValue:101 time:4 coord:1}
{rawValue:137 time:5 coord:12}
{rawValue:173 time:6 coord:23}
{rawValue:191 time:7 coord:16}
{rawValue:213 time:8 coord:13}
{rawValue:245 time:9 coord:20}
{rawValue:259 time:10 coord:9}
{rawValue:277 time:11 coord:2}
{rawValue:311 time:12 coord:11}
{rawValue:347 time:13 coord:22}
{rawValue:369 time:14 coord:19}
{rawValue:385 time:15 coord:10}
{rawValue:403 time:16 coord:3}
{rawValue:431 time:17 coord:6}
{rawValue:467 time:18 coord:17}
{rawValue:499 time:19 coord:24}
{rawValue:515 time:20 coord:15}
{rawValue:529 time:21 coord:4}
{rawValue:557 time:22 coord:7}
{rawValue:593 time:23 coord:18}
{rawValue:621 time:24 coord:21}
```
