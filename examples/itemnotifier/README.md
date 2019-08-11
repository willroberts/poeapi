#### itemnotifier

This example searches in real time, until the user exits with Ctrl-C, for Kaom's
Heart in Standard league. When one is listed for sale, it prints the character
name and asking price (if there is one).

This is similar to the live search features of trade websites, but on the
command line instead. This code could be extended for a number of purposes, such
as having the program play a sound alert ("WOOP!"), displaying a popup
notification, sending an email, etc.

Output:

```
$ go run main.go
2019/08/11 15:53:06 LowFattt is selling Kaom's Heart in Standard league.
2019/08/11 15:53:08 StriXXa is selling Kaom's Heart for ~price 75 chaos in Standard league.
2019/08/11 15:53:22 SASHQUE is selling Kaom's Heart in Standard league.
2019/08/11 15:53:28 preXisTence is selling Kaom's Heart for ~b/o 22 exa in Standard league.
2019/08/11 15:54:04 solonaras is selling Kaom's Heart in Standard league.
2019/08/11 15:54:24 EvilLoki is selling Kaom's Heart for ~price 1.3 exa in Standard league.
2019/08/11 15:54:27 PaladinLT is selling Kaom's Heart for ~price 1.3 exa in Standard league.
```