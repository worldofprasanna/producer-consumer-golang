# Producer Consumer problem

## Statement
Producer should produce the items, consumer should consume it. Have a buffer of 10 items (so only after producer producing 10 items, consumer should start its work). When sigterm call is given, consume the remaining items and exit the process. When sigkill call is given, terminate the process immediately without consuming the remaining items. If no signal is given, stop the process after 30 seconds.

## Topics to cover
- Channels
- Handle Signals, Graceful termination
- Range, Switch
- Wait Group
