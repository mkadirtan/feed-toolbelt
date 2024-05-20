Set FontSize 30
Output example.gif
Set Width 1280
Set Height 600

Type "feed-toolbelt find nooptoday.com"
Sleep 150ms
Enter
Sleep 900ms

Type "echo nooptoday.com | feed-toolbelt find"
Sleep 150ms
Enter
Sleep 900ms

Type "curl -s https://nooptoday.com | feed-toolbelt find --pipe"
Sleep 150ms
Enter
Sleep 900ms

Type "feed-toolbelt find nooptoday.com --strategy-common"
Enter
Sleep 9500ms