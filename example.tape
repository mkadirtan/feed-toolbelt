Set FontSize 30
Set TypingSpeed 85ms
Output example.gif
Set Width 1280
Set Height 600

Type "feed-toolbelt find nooptoday.com"
Sleep 250ms
Enter
Sleep 1200ms

Type "echo nooptoday.com | feed-toolbelt find"
Sleep 250ms
Enter
Sleep 1200ms

Type@105ms "curl -s https://nooptoday.com | feed-toolbelt find --pipe"
Sleep 250ms
Enter
Sleep 1200ms

Type "feed-toolbelt find nooptoday.com --strategy-common"
Sleep 150ms
Enter
Sleep 9500ms