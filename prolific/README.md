The files are intended for use with [prolific](https://github.com/onsi/prolific), and should then be imported into Pivotal Tracker projects for each person to be onboarded.

Import them in reverse order - newly-imported stories will go to the top of the icebox.

Each story requires a label to show who it is meant for. Suggested labels are:

- All-roles
- Engineer
- Product-manager

Adding these to the relevant prolific stories means that when they are added to tracker you can filter by them, and just copy out the stories you need for the person you are about to onboard.

In terms of overall structure, the programme adds files (e.g linux.prolific, golang.prolific) in reverse alphabetical order. We can specify the order they go in by labelling each file. EG:
015 - general.prolific
.
007 - git.prolific
.
000 -final.prolific
In the tracker, this would place the contents of final.prolific LAST in the icebox, and general.prolific FIRST.
Suggest we add another story that asks us to write something pulling in files in reverse order, making this naming convention easier (starting at 001 and going upwards).