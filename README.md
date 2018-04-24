[![Codacy Badge](https://api.codacy.com/project/badge/Grade/006a742933a14655b987448cb7e02f5e)](https://app.codacy.com/app/alexdomoryonok/AWS-S3-Slack-Lambda-Notification?utm_source=github.com&utm_medium=referral&utm_content=Domoryonok/AWS-S3-Slack-Lambda-Notification&utm_campaign=badger)


Lambda function for notification about items which uploaded to some s3 bucket


```
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNmmdhhyyysssssyyyyhhhdmNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMmdhso+//:::::::::::::::::::::::::/+oyhmNMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMmhs+/:::::::::::::::::::::::::::::::::::::::/+sdNMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMNho/::::::::::::::::::::::::::::::::::/+ooooo+/:::::/ohNMMMmhyssyhmNMMMMMMMMMM
MMMMMMMMNhysoosyhmMMho/:::/+ooo++++ooo/:::::::::::::::::/oo+/-.````.:+os/::::/sds/:::::::/smMMMMMMMM
MMMMMMdo/:::::::/yy/:::/oo/-```   ```./oo/::::::::::::os/``            `/s+::::/oy+:/::::::/yMMMMMMM
MMMMNo:::::/:::sy/:::/s/.               .+s/::::::::/y:`                 `oo:::::/ymNmy::::::yMMMMMM
MMMM+:::/ymNmdho::::oo`                   .y/::::::/y.                     oo::::::oNMMs:::::/NMMMMM
MMMd::::sMMMMd/::::+s   .--.               .y::::::h. `/shy+.              `h:::::::+mm/::::::NMMMMM
MMMd::::+dMMd/:::::h` .ymNNmy.              s+::::/y  yMMMMMm.              s/:::::::+h::::::oMMMMMM
MMMMo::::/om/::::::d  hMMMMdNy              oo::::+s `NMMMh/N:              s+::::::::ss:::/yNMMMMMM
MMMMNy+:::h+:::::::d` sMMMN:ys              s+:::::h` +NMMmyy`             `h::::::::::d+ohmMMMMMMMM
MMMMMMNdysy::::::::o+ `+dmmh+`             .y::::::+s` ./++-              `s+::::::::::oNMMMMMMMMMMM
MMMMMMMMMN/:::::::::s/  ```               -y/:::///:+s-                 `:s+::::::::::::dMMMMMMMMMMM
MMMMMMMMMh:::::::::::oo-                .+o//oydmmmdhsso:.`          `-/oo/:::::::::::::oMMMMMMMMMMM
MMMMMMMMMo::::::::::::/oo:.``      `.-/oo/:/dMMMMMMMMMd:+oo+/:---:/+oo+/::::::::::::::::/NMMMMMMMMMM
MMMMMMMMN/:::::::::::::::/+o++++++++++/::/osNMMMMMMMMNho+/:://+++///:::::::::::::::::::::dMMMMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::/s+-.:oyhhhys+:.-/oo/:::::::::::::::::::::::::::::sMMMMMMMMMM
MMMMMMMMd::::::::::::::::::::::::::::::y:.................:y:::::::::::::::::::::::::::::+MMMMMMMMMM
MMMMMMMMh::::::::::::::::::::::::::::::y:.....--::---.....-y/::::::::::::::::::::::::::::/MMMMMMMMMM
MMMMMMMMh::::::::::::::::::::::::::::::/so++++/:oh-:/+o//+s+::::::::::::::::::::::::::::::NMMMMMMMMM
MMMMMMMMh:::::::::::::::::::::::::::::::::o+    +s    d///::::::::::::::::::::::::::::::::mMMMMMMMMM
MMMMMMMMd:::::::::::::::::::::::::::::::::oo    o+    d:::::::::::::::::::::::::::::::::::dMMMMMMMMM
MMMMMMMMd::::::::::::::::::::::::::::::::::h.  `ys    d:::::::::::::::::::::::::::::::::::dMMMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::::::/soos+os+os+:::::::::::::::::::::::::::::::::::hMMMMMMMMM
MMMMMMMMN/::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::hMMMMMMMMM
MMMMMMMMM+::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::hMMMMMMMMM
MMMMMMMMMo::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::hMMMMMMMMM
MMMMMMMMMy::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::yMMMMMMMMM
MMMMMMMMMd::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::yMMMMMMMMM
MMMMMMMMMm::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::yMMMMMMMMM
MMMMMMMMMM::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::yMMMMMMMMM
MMMMMMMMMM::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::sMMMMMMMMM
MMMMMMMMMM/:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::oMMMMMMMMM
MMMMmdyysd+:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::oy+oshmMMM
MNs/--...y+:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::+y....-/dM
N::-.....h+:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::/h....-/:N
Mho-..-+sN/::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::ms/-.-/dM
MMyoshmMMM:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::mMMmdhdMM
MMMMMMMMMN:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::dMMMMMMMM
MMMMMMMMMd:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::hMMMMMMMM
MMMMMMMMMh:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::yMMMMMMMM
MMMMMMMMMs:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::sMMMMMMMM
MMMMMMMMMo:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::oMMMMMMMM
MMMMMMMMM/:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::+MMMMMMMM
MMMMMMMMN/:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::+MMMMMMMM
MMMMMMMMN::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::+MMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::oMMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::sMMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::hMMMMMMMM
MMMMMMMMm::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::dMMMMMMMM


```
