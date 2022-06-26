Cbz-merger was developed as a simple project to help you when in need of tool to merge .cbz files.

- How to use it?
  Just extract the .cbz files you want to merge to a separated folder, keep the files of each chapter inside its
  respective folder. Then in a terminal window execute the app passing the path to the root folder where you extracted
  the chapters and the name you want to give to the .cbz file that will be created. Example:
  ./cbz-merger merge "/home/{user}/.../Berserk/Berserk 21" "Berserk 21"
  In the line above we first execute the file passing the "merge" flag then the path to the folder where we first extracted the files
  then the name we wish our result file will have, so the result file will be named as "Berserk 21".
  PS.: Take in consideration that this is an simple application, builted to help me with some .cbz files i wished to merge, and I decided to share
  with anyone that its having problems with this kind of task.
  PS2.: More functionalities will be added in the future, like extracting automatically the .cbz files, but you can use any program that suports this kind of file,
  I used PeaZip.
  PS3.: This program was developed in golang 1.18.1 and to run on Fedora 36 Linux dist, and was not tested in other OS like Ubuntu or Windows, so take this
  consideration.
