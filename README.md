d20 simulator
=============

This simple program's intent is to see what a house rule's consequences has on d20 ability check and attack rolls in D&D fifth edition.

In the house rule, the player can re-roll their d20 if they roll at or less than their modifier.  For instance, if a player has a +5 attack bonus on a roll, and rolls 5 or less, they can re-roll their die.  But only once.

Why?  The d20 system creates evenness in chances for all 20 outcomes on the d20, yielding a feeling of one's proficiency and modifiers meaning less than they should.  This house rule should overcome this issue.

Graph
-----

The program yields a program of modifiers +0 through +6 from DCs 5 through 22:

X-axis is the percentage of hits (all the way at the top text is 100%).
Y-axis is the DCs tried.

    +0: +0    * = standard, # = new




     *
         *   *
             #   #
                 *   *
                         *
                             *
                                 *
                                     *   *
                                             #
                                             *   *
                                                     *
                                                         *
                                                             *
     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +1: +1    * = standard, # = new


     #
     *   #
         *   #
             *   #
                 *   #
                     *   *
                             *
                                 *   #
                                     *
                                         *   #
                                             *   *
                                                     *
                                                         *
                                                             *
                                                                 *
     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +2: +2    * = standard, # = new

     #
     *   #
         *   #   #
             *       #
                 *       #
                     *       #
                         *   *   #
                                 *   #
                                     *   #
                                         *   #
                                             *   #
                                                 *   #
                                                     *   *
                                                             *
                                                                 *   #
                                                                     *
     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +3: +3    * = standard, # = new

     *   #   #
         *       #
             *       #
                 *       #
                     *   *   #
                                 #
                             *   *   #
                                     *   #
                                         *   #   #
                                             *
                                                 *   #
                                                     *   *
                                                             *
                                                                 *
                                                                     *   #
                                                                         *
     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +4: +4    * = standard, # = new
     *
         *   #   #   #
             *           #
                 *           #
                     *           #
                         *           #
                             *   *       #
                                     *       #
                                         *       #
                                             *       #
                                                 *       #
                                                     *   *   #
                                                             *   #
                                                                 *   #
                                                                     *   #
                                                                         *

     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +5: +5    * = standard, # = new
     *   *
             *   #   #   #
                 *           #
                     *           #
                         *           #
                             *           #
                                 *   *       #
                                         *       #
                                             *       #
                                                 *       #
                                                     *       #
                                                         *       #
                                                             *   *   #
                                                                     *   #
                                                                         *


     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22


    +6: +6    * = standard, # = new
     *   *   *
                 *   #   #   #
                     *           #   #
                         *               #
                             *               #
                                 *               #
                                     *   *           #
                                             *
                                                 *       #
                                                     *       #
                                                         *       #
                                                             *       #
                                                                 *   *   #
                                                                         *



     5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22