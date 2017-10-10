d20 simulator
=============

This simple program's intent is to see what a house rule's consequences has on d20 ability check and attack rolls in D&D fifth edition.

In the house rule, the player can re-roll their d20 if they roll at or less than their modifier.  For instance, if a player has a +5 attack bonus on a roll, and rolls 5 or less, they can re-roll their die.  But only once.

Why?  The d20 system creates evenness in chances for all 20 outcomes on the d20, yielding a feeling of one's proficiency and modifiers meaning less than they should.  This house rule should overcome this issue.

Graph
-----

The program yields a program of modifiers +0 through +10 with varies DCs, beginning with its modifier +1 and 17 more.  So +0 gets 1 through 18, +1 gets 2 through 19, and so on.

X-axis is the percentage of hits (all the way at the top text is 100%).
Y-axis is the DCs tried.

    +0    * = standard, # = new, @ = both

    100%  @                                                                      
                                                                                 
              @                                                                  
                  @                                                              
                      @                                                          
                          @                                                      
                                                                                 
                              @                                                  
                                  @                                              
                                      @                                          
                                                                                 
                                          @                                      
                                              @                                  
                                                  @                              
                                                      @                          
                                                                                 
                                                          @                      
                                                              @                  
                                                                  @              
                                                                      @          
                                                                          #      
                                                                          *      
                                                                              @  
                                                                                 
      0%                                                                         

     DC:  1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  


    +1    * = standard, # = new, @ = both

    100%  @                                                                      
              #                                                                  
              *   #                                                              
                  *   #                                                          
                      *   #                                                      
                          *                                                      
                              #                                                  
                              *   #                                              
                                  *   #                                          
                                      *                                          
                                          #                                      
                                          *   #                                  
                                              *   #                              
                                                  *                              
                                                      @                          
                                                          @                      
                                                              #                  
                                                              *                  
                                                                  @              
                                                                      @          
                                                                          #      
                                                                          *      
                                                                              @  
                                                                                 
      0%                                                                         

     DC:  2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  


    +2    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #                                                              
              *       #                                                          
                  *                                                              
                      *   #                                                      
                              #                                                  
                          *       #                                              
                              *                                                  
                                  *   #                                          
                                      *   #                                      
                                          *   #                                  
                                                                                 
                                              *   #                              
                                                  *   #                          
                                                      *                          
                                                          #                      
                                                          *   #                  
                                                              *   #              
                                                                  *              
                                                                      @          
                                                                          @      
                                                                              #  
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  


    +3    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #   #                                                          
              *                                                                  
                  *       #                                                      
                      *       #                                                  
                          *       #                                              
                                                                                 
                              *       #                                          
                                  *       #                                      
                                      *                                          
                                          *   #                                  
                                                  #                              
                                              *                                  
                                                  *   #                          
                                                      *   #                      
                                                              #                  
                                                          *                      
                                                              *   #              
                                                                  *   #          
                                                                      *          
                                                                          @      
                                                                              #  
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  


    +4    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #   #   #                                                      
              *                                                                  
                  *           #                                                  
                      *           #                                              
                          *                                                      
                                      #                                          
                              *           #                                      
                                  *                                              
                                      *       #                                  
                                                                                 
                                          *       #                              
                                              *       #                          
                                                  *       #                      
                                                      *                          
                                                          *   #                  
                                                                                 
                                                              *   #              
                                                                  *   #          
                                                                      *   #      
                                                                          *      
                                                                              #  
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  


    +5    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #   #                                                          
              *           #   #                                                  
                  *                                                              
                      *           #                                              
                                      #                                          
                          *                                                      
                              *           #                                      
                                  *           #                                  
                                      *                                          
                                                  #                              
                                          *           #                          
                                              *                                  
                                                  *       #                      
                                                      *                          
                                                          *   #                  
                                                                  #              
                                                              *                  
                                                                  *   #          
                                                                      *   #      
                                                                                 
                                                                          *   #  
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  


    +6    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #                                                              
              *       #   #   #                                                  
                  *               #                                              
                      *               #                                          
                                                                                 
                          *               #                                      
                              *                                                  
                                  *           #                                  
                                      *           #                              
                                                                                 
                                          *           #                          
                                              *                                  
                                                  *       #                      
                                                      *       #                  
                                                          *                      
                                                                  #              
                                                              *       #          
                                                                  *              
                                                                      *   #      
                                                                                 
                                                                          *   #  
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  


    +7    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #                                                              
              *       #   #                                                      
                  *           #   #                                              
                      *               #                                          
                                          #                                      
                          *                                                      
                              *               #                                  
                                  *                                              
                                      *           #                              
                                                      #                          
                                          *                                      
                                              *           #                      
                                                  *                              
                                                      *       #                  
                                                          *       #              
                                                                                 
                                                              *       #          
                                                                  *              
                                                                      *   #      
                                                                          *   #  
                                                                                 
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  


    +8    * = standard, # = new, @ = both

    100%  @                                                                      
              #   #                                                              
              *       #                                                          
                  *       #   #                                                  
                      *           #   #                                          
                                          #                                      
                          *                   #                                  
                              *                                                  
                                  *               #                              
                                      *                                          
                                          *           #                          
                                                          #                      
                                              *                                  
                                                  *           #                  
                                                      *                          
                                                          *       #              
                                                                                 
                                                              *       #          
                                                                  *              
                                                                      *   #      
                                                                          *   #  
                                                                                 
                                                                              *  
                                                                                 
      0%                                                                         

     DC:  9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  


    +9    * = standard, # = new, @ = both

    100%  @                                                                      
              #                                                                  
              *   #   #                                                          
                  *       #   #                                                  
                      *           #   #                                          
                          *               #                                      
                                              #                                  
                              *                   #                              
                                  *                                              
                                      *               #                          
                                                                                 
                                          *               #                      
                                              *                                  
                                                  *           #                  
                                                      *                          
                                                          *       #              
                                                                      #          
                                                              *                  
                                                                  *       #      
                                                                      *          
                                                                          *   #  
                                                                                 
                                                                              *  
                                                                                 
      0%                                                                         

     DC: 10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  


    +10    * = standard, # = new, @ = both

    100%  @                                                                      
              #                                                                  
              *   #   #                                                          
                  *       #                                                      
                      *       #   #                                              
                          *           #   #                                      
                                              #                                  
                              *                   #                              
                                  *                                              
                                      *               #                          
                                          *                                      
                                                          #                      
                                              *               #                  
                                                  *                              
                                                      *           #              
                                                          *                      
                                                                      #          
                                                              *                  
                                                                  *       #      
                                                                      *          
                                                                          *   #  
                                                                                 
                                                                              *  
                                                                                 
      0%                                                                         

     DC: 11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  