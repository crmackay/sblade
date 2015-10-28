package threeprime

// test data. Here we will use the "misc" third line in the FASTQ reads to store our
//expected test results. Namely, what the 3'-cut read should look like
var testRawReads = `@HWI-ST560:155:C574EACXX:3:1101:2012:1985 1:N:0:
AGCAGGGAGGACGATGCGGTTGTGATATAATACAACCTGCTAAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGTTGTGATATAATACAACCTGCTAA
@CCFFFFFHHHHHIIIIIICDF?D@BBD<<??FHG?GIIIIEG@@A@CAFHIIEIEEH@AHFF9AB;?CCCA@3=A<ACA>@CCB@::ABBBBBBCBBBB
@HWI-ST560:155:C574EACXX:3:1101:2168:1987 1:N:0:
AGCAGGGAGGACGATGCGGTACAGATCACATTGCCAGGGATTACCACAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGTACAGATCACATTGCCAGGGATTACCACA
@CCFFFFFHHHGHIJJJJJJJGJJFJGDHIJJIIJJJIIGGIGEHIJIJJGIJJGHAHCHEEEFFBC:B?DBDCCA@>BDDBCD@ACCABDCCBB#####
@HWI-ST560:155:C574EACXX:3:1101:2074:1991 1:N:0:
GTGAGGGAGGACGATGCGGCTGAGTAGCTTATCAGACTGATGTTGACGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAA
+GTGAGGGAGGACGATGCGGCTGAGTAGCTTATCAGACTGATGTTGAC
BCCFFFFFHHHHHJJJJJJIJJJIDHIJJIGGGJIIJJJJIJIJFJJJJJJHGHHHHFFFFFFDECDDDDDDDDDCBBDDDDDDDDDDACDDDDDDDDDD
@HWI-ST560:155:C574EACXX:3:1101:2214:1998 1:N:0:
AGCAGGGAGGACGATGCGGAACTGATGTCTAAGTACGCACGGCCGGTACAGTGAAACTGCGAATGGCTCGTGTCAGTCACTTCCAGCGGGCGTATGCCGT
+AGCAGGGAGGACGATGCGGAACTGATGTCTAAGTACGCACGGCCGGTACAGTGAAACTGCGAATGGCTC
CCCFFFFFHHHHHJJJJJJJJJJJJJJJJIIGHEIIJIGIGGIGFFCCDDECCDDDDDDDDDDDDCDCC-9?28A@>4@AAC3>4439B###########
@HWI-ST560:155:C574EACXX:3:1101:2300:1939 1:N:0:
CACAGGGAGGACGATGCGGGAGTGAGACCGTCTTGCTTACTTGTCCGATGAAATGAATGAAATAGAAAGTGGGAAAATAATGTGTCAGTCACTTCCAGCG
+CACAGGGAGGACGATGCGGGAGTGAGACCGTCTTGCTTACTTGTCCGATGAAATGAATGAAATAGAAAGTGGGAAAATAAT
BBCFFFFFHHHHHJJJJJJJIJHIGIHHIJJJJJIGIIEHIIGHHHHFFFCEDDEDCCCDDDACDCDDD@CDDACBCCCDDCDEDDCDCCCCCDCC:@99
@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:
CACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGGGT
+CACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGG
BBBFFFFFHHHHHFIJIJJJJJJJJIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD#####
@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:
AGCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTG
BCCFFFFFHHHHHJJJJJJJIJJJGJIJJJIJJJJJJJJJJGJIJIJHHHGFFFFFFEEEEEEEDDDBDDDDDDDCDDDDDDDDDDDDDDDDDDD@B@DB
@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:
AGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCG
BCCFFFFFHHHHHIJJJJJJJJJIIHIIJJJFHJJJJJJJJJJJJJJHHFFFFFFEEEEEEEDDDDDDDDDDCDDBBBDDDDDDDDCDDDDDDD9BB>BD
@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:
GCTAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGCGTGTCAGTCACTTCCAGCGGGTGTCAGTCACTTCCAGCGGTCGTATG
+GCTAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGC
@@@FFFFFHHGHHJJJJJJGIEFHFHGDHGIEGGHIIJIICHHIJHEFHGDDDCDD@BCCDDDDDDA@CDDDDD@><ACDDCCCCCC>CC?>B9@B>833
@HWI-ST560:155:C574EACXX:3:1101:2425:1982 1:N:0:
GTGAGGGAGGACGATGCGGTTGTGTGAGAACTGAATTCCATAGGCTGTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA
+GTGAGGGAGGACGATGCGGTTGTGTGAGAACTGAATTCCATAGGCTGT
CCCFFFFFHHFHHIJJJJJIJIGIIJJJFGIIIIIJJJG<FHEHIJIGIHGHJFIHHHHHHFFFFEDDDB@=?CD@ABBBDDDDCCDDC@DDDDDDB@BB
@HWI-ST560:155:C574EACXX:3:1101:2309:1987 1:N:0:
GCTAGGGAGGACGATGCGGAGAAGAGAAGGAAGATGCTAATGGCAATATCGTCTATGAGAAGAACTGCGTGTCAGTCACTTCCAGCGGTCGTATGCCGTC
+GCTAGGGAGGACGATGCGGAGAAGAGAAGGAAGATGCTAATGGCAATATCGTCTATGAGAAGAACTGC
CCCFFFFFHHHHHJJJJJJGIHJJJJIJIJJJJIIJJJJJIIJJJIIJJIIHHHHHFFFFFFEEEEEDDDDDDDDDDDDDDDDDDDDDBDDDBDDD>959
@HWI-ST560:155:C574EACXX:3:1101:2364:1995 1:N:0:
AGCAGGGAGGACGATGCGGAAACGTAGTGTTTCCTACTTTATGGATGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGAAACGTAGTGTTTCCTACTTTATGGAT
@CCFFFFFHHHHHIJJJJJJJJJJGHJIIIIJJJJIJJJJHIJIIIJJJJJJJJHHGHHHFFFFDDDDDDDDDDDBDDDDDDDDDDDDDDDDDBD@DD9B
@HWI-ST560:155:C574EACXX:3:1101:2342:1996 1:N:0:
GTGAGGGAGGACGATGCGGCGGGGTGGCTGTTACTTCCAGCGGTGGGTGTCCGTCTTTTGCTTGGAAAAAATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+GTGAGGGAGGACGATGCGGCGGGGTGGCTGTTACTTCCAGCGGTGGGTGTCCGTCTTTTGCTTGGAAAAAATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
@@@DDDDDHHHHHIIIIII>6;6?############################################################################
@HWI-ST560:155:C574EACXX:3:1101:2654:1937 1:N:0:
NCTAGGGAGGACGATGCGGAGAAGAGTGAATGGCAACAATGAAGCTATCGTGCATGTTGTGGAGACTCGTGTCAGTCACTTCCAGCGGTCGTATGCCGTC
+NCTAGGGAGGACGATGCGGAGAAGAGTGAATGGCAACAATGAAGCTATCGTGCATGTTGTGGAGACTC
#1=DDFFFHHHHHJJJJJJIIIJIIIHIGIIJJIIJIIJJIJJJIIJJIJIHHGHHFFFFFDEDEDDDDDDDDDDDDDDDDDCCDDDDBDDDDDDCC@99
@HWI-ST560:155:C574EACXX:3:1101:2588:1943 1:N:0:
GCTAGGGAGGACGATGCGGTAGAGTGTAGTGTTTCCTACTTTATGGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+GCTAGGGAGGACGATGCGGTAGAGTGTAGTGTTTCCTACTTTATGG
@CCFFFFFGHHFFHIIIIIIGEII?FD@GGGGIIIGHHIIII@FHGICGGCFHCHGGGIHHHG@<BACD?<ACCDBB=B@CAA@CDC?ACCAB#######
@HWI-ST560:155:C574EACXX:3:1101:2514:1947 1:N:0:
CACAGGGAGGACGATGCGGATATGAGAATGTGTGGTAAATTGAATAAAGCTAGCCGTGATCCTCAGCTGTTGCTGCGTGTCAGTCACTTCCAGCGGTCGT
+CACAGGGAGGACGATGCGGATATGAGAATGTGTGGTAAATTGAATAAAGCTAGCCGTGATCCTCAGCTGTTGCTGC
@BBFFFFFHHHHHIJJJJJJJJJJIIIJIJHIFHIFGGIJJJJIIJJIIGGIJJJJHHFFFFFFEEEEEEDDDDDDDDDDDDDDDDDDDDDDDDDDD@D9
@HWI-ST560:155:C574EACXX:3:1101:2541:1953 1:N:0:
GTGAGGGAGGACGATGCGGTAAAGTAGCTTATCAGACTGATGTTGACGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAA
+GTGAGGGAGGACGATGCGGTAAAGTAGCTTATCAGACTGATGTTGAC
@CCDFFFFHHGHHIJJJJJJJJJJFHJJJJIJJJJJJJJJJJJJJJJJJJJJJIJJJJHHHHHHEFDDDDDDDDDDDDDDDDDDDDDCBDDDDDDDBBDD
@HWI-ST560:155:C574EACXX:3:1101:2722:1957 1:N:0:
GCTAGGGAGGACGATGCGGGTTAGATGACCCGCCGGGCAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGCGTGTCAGTCACTTC
+GCTAGGGAGGACGATGCGGGTTAGATGACCCGCCGGGCAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGC
CCCFFFFFHHHHHJJJJJJJIIJIJJJJJJJJJJJJHFFFDEEEEDDDDDDDDDDDDDDEEDDDDDDDDDDDDDDDBBCDEDDDDDDDDDDDDDD>::CC
@HWI-ST560:155:C574EACXX:3:1101:2526:1992 1:N:0:
AGCAGGGAGGACGATGCGGCATGGATGACCTCTCTGCCACACTGCAATTTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAA
+AGCAGGGAGGACGATGCGGCATGGATGACCTCTCTGCCACACTGCAATTT
@CCFFFFFHHHHHJJJJJJJJJJJJJIJJJJIJIIIIJGIGIJJJJIIJJGHGHHGFFFFFFFEEDEEDDBDBDDCDDDDDDDDDDDCDDCADCDDDBD0
@HWI-ST560:155:C574EACXX:3:1101:2618:1997 1:N:0:
GCTAGGGAGGACGATGCGGAATCGGGGCAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGCGTGTCAGTCACTTCCAGCGGTCGT
+GCTAGGGAGGACGATGCGGAATCGGGGCAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGC
BC@FFFFFHHHHHJJIJJJJJJJJJJJJJJJIIHIHJHFFFFFDDDEEDDDDDDDDDDDDDDDDD;BBCCDDDDDDDDDDDDDDDDDDDACCDCDBB@BB
@HWI-ST560:155:C574EACXX:3:1101:2867:1949 1:N:0:
AGCAGGGAGGACGATGCGGAAGTGGTGTCAGTCACTTCCAGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGAAGTG
B@CFFFFFHHHHHJJJJJJIJIEHFIGIJJJJJJJJJIIJIJIIIJHHFHHHFFFFFFEEDDDBDDDDDDDDDDDABDDDDDDDDDDDDD@55><9@B<>
@HWI-ST560:155:C574EACXX:3:1101:2792:1953 1:N:0:
CACAGGGAGGACGATGCGGAAAAGTTGTCCAATATGAGAACGGTAGCTTGAGGAGTAGAGACGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGC
+CACAGGGAGGACGATGCGGAAAAGTTGTCCAATATGAGAACGGTAGCTTGAGGAGTAGAGAC
BCCFFFFFHHHHHIJJJJJJJJJJHGIJJJJJJJJIIIJJJJJHIJJJJJJHGHHFDDFFFEDDDDDDDDDDDDDDDDDDDDBDDDBDDDDDBDDDCDCA
@HWI-ST560:155:C574EACXX:3:1101:2988:1954 1:N:0:
CACAGGGAGGACGATGCGGACACGAAAGAAGGAGCTGGACCATGAAGATAGCCACTTGAACTTGGATGAGACAGCCAGGCTCCTGCGTGTCAGTCACTTC
+CACAGGGAGGACGATGCGGACACGAAAGAAGGAGCTGGACCATGAAGATAGCCACTTGAACTTGGATGAGACAGCCAGGCTCCTGC
BCCFFFFFHHHHHJJJJJJJJJJJJJJJJJJJIJJJIJJJJJJHHHHHFFFFFFEEEEEEDDDDDDDCDDDDDDDABDDBDDDDDDD<BDDDDDDC@@CD
@HWI-ST560:155:C574EACXX:3:1101:2936:1962 1:N:0:
CACAGGGAGGACGATGCGGCAACGAGGACGATGCGGCGGCGAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGCGTGTCAGTCAC
+CACAGGGAGGACGATGCGGCAACGAGGACGATGCGGCGGCGAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGC
BBBFFFFFHHHHHJJJJIJJJJJIHJJJIIIJJIJHHDDDBBDDDDDDDDDDDBDDDDDDCCDEDCDDDDDDDCDDBD@D?CDEC@CDDDDDDDACCC:4
@HWI-ST560:155:C574EACXX:3:1101:2875:1972 1:N:0:
CACAGGGAGGACGATGCGGTTGAGAGGAAGTCTTATTTGTCAAAAGCAGTACATAGAGCATGTGCTTGTGGCGTGTCAGTCACTTCCAGCGGTCGTATGC
+CACAGGGAGGACGATGCGGTTGAGAGGAAGTCTTATTTGTCAAAAGCAGTACATAGAGCATGTGCTTGTGGC
@?@FFFFFHHHHHJJJJJJIJJGI=FH@FEDDFFGIJIGDBGGGIJEFHHJIEHG<GEHFEHFFFDFBFEACBD@?<@:>CD:@A:3>4(5><9@#####
@HWI-ST560:155:C574EACXX:3:1101:2904:1980 1:N:0:
CACAGGGAGGACGATGCGGTAAAGAGAAAGATCCGTATGTGTGTTGTTTACAAAGGTTGTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTT
+CACAGGGAGGACGATGCGGTAAAGAGAAAGATCCGTATGTGTGTTGTTTACAAAGGTTGT
@@@FFFFFHHHHHJJJJJJHGIEH>A??F@?@DBGD?FH<BGEHHHHHE@@4CC?A?C;B;?BE>@@CC>@@A>@;5,388>90<CC::5)0?B>@:+>@
@HWI-ST560:155:C574EACXX:3:1101:2990:1980 1:N:0:
GTTGATGCTAAAGGTGAGCCGCTTAAAGCTACCAGTTATATGGCTGTTGGTTTCTATGTGGCTAAATACGTTAACAAAAAGTCAGATATGGACCTTGCTG
+GTTGATGCTAAAGGTGAGCCGCTTAAAGCTACCAGTTATATGGCTGTTGGTTTCTATGTGGCTAAATACGTTAACAAAAAGTCAGATATGGACCTTGCTG
<@@FDFFFHHHHHJEHHJJJJJJJJJJJJJJJJJIJJIEIJIJIEGIJFIHHIHIGIIAHDGIDEHHGHHHHFDDBDFDDD@CDDDDDDDCDCCDCACDD
@HWI-ST560:155:C574EACXX:3:1101:2923:1989 1:N:0:
CACAGGGAGGACGATGCGGAGATGTTGCATGATGACTTGAATTGTCGGATACCCCTCACCCCGTTCATGGGTGAGAAACAGCTAGTCTGACGTGTCAGTC
+CACAGGGAGGACGATGCGGAGATGTTGCATGATGACTTGAATTGTCGGATACCCCTCACCCCGTTCATGGGTGAGAAACAGCTAGTCTGAC
BCCFFFFFHHHHHJJJJJJJJIJJIHIJJJIJJJIJJJJIJJJHJJJJIJHHHGHFFFDEEDDDDDDECDDCDDDDCC?CC<CCCCDEDDDDDDDD:A::
@HWI-ST560:155:C574EACXX:3:1101:2969:1991 1:N:0:
GTTGATGCTAAAGGTGAGCCGCTTAAAGCTACCAGTTATATGGCTGTTGGTTTCTATGTGGCTAAATACGTTAACAAAAAGTCAGATATGGACCTTGCTG
+GTTGATGCTAAAGGTGAGCCGCTTAAAGCTACCAGTTATATGGCTGTTGGTTTCTATGTGGCTAAATACGTTAACAAAAAGTCAGATATGGACCTTGCTG
@BCFFFFFHHHHHJCEEHIJJJJJJJJIJJJJJJIJJGIIJIIJJIJIJJJJJIIGJEGIIJIIGIIIHHHHGFFEFFDDDCCEDDDDDDDDDDDDDDDC
@HWI-ST560:155:C574EACXX:3:1101:2877:1994 1:N:0:
AGCAGGGAGGACGATGCGGACAAGTGAGAACTGAATTCCATAGGCTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGACAAGTGAGAACTGAATTCCATAGGCT
@CCFFFFFHHGHHIIIIIIEGECHGH<B@@BFFHGGG>CFHGE>CFHH@@@G@C>CEHHCEC>?AC@B?;A?A@C>59A@>@44@C42@CDDDDDDD<BB
@HWI-ST560:155:C574EACXX:3:1101:3086:1952 1:N:0:
CACAGGGAGGACGATGCGGTTAAGAGCCTGAATAATACTGTGAACTGCACATGCACGCCCGGGCGTGAGCTGGTGCTTGTGCTCGCTCGTGTCAGTCACT
+CACAGGGAGGACGATGCGGTTAAGAGCCTGAATAATACTGTGAACTGCACATGCACGCCCGGGCGTGAGCTGGTGCTTGTGCTCGCTC
BBBFFFFFHHHHHJJJJJJJJJJIJJIJJJJJJJJIJJJIJJJIJJJJJJJJIJJJIHHHFFDDDDDDDDDDDDDACDDDDDDDDDDDDDDDDCCCC44:
@HWI-ST560:155:C574EACXX:3:1101:3206:1953 1:N:0:
GCTAGGGAGGACGATGCGGCACAGGTTTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGCGTGTGTGTATTAAACCCACTGAACGTGACAGA
+GCTAGGGAGGACGATGCGGCACAGGTTTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGCGTGTGTGTATTAAACCCACTGAAC
CCCFFFFFHHHHHIJJJJJJJJHIJGGHHHGHHCFBFFHEHHGHHIHHHHFFFFEEEEEBBDDDD?B(;@DDDD8?:4+:4>A(8<8((444?B######
@HWI-ST560:155:C574EACXX:3:1101:3018:1971 1:N:0:
AGCAGGGAGGACGATGCGGTTAAGGTGTCAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGTTAAG
BCCFFFFFHHHHHIJJJJJJJJJJJIJJJJJJJJJJJIJJJJJJJJJJJJJHHFFFEEEDBDDDDDCDDDDBCDD<@#######################
@HWI-ST560:155:C574EACXX:3:1101:3156:1972 1:N:0:
GTGAGGGAGGACGATGCGGTTGTGAATAATCTGTTTCGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAAAAAAAA
+GTGAGGGAGGACGATGCGGTTGTGAATAATCTGTTTC
@CCFFFFFHHHHHIJJJJJJJJHIJJIJJJIJJJJJJIJIJJJJJJJIJJJJJIIHFFFDDDDDDDDDDDDDDDDCDDDDDDDBB><559955*5<BDB@
@HWI-ST560:155:C574EACXX:3:1101:3075:1975 1:N:0:
CACAGGGAGGACGATGCGGTATGGATGGCAAAGACCTGGCTTCTGTGAACAACTTGCTGAAAAAGCATCAGCTGCTAGAGGCAGACGTGTCAGTCACTTC
+CACAGGGAGGACGATGCGGTATGGATGGCAAAGACCTGGCTTCTGTGAACAACTTGCTGAAAAAGCATCAGCTGCTAGAGGCAGAC
BBBFFFFFHHHFDHIIHIIJDGGIJIACCGGIBD?F??9BFFHI<GDGGIIGIDEECEHHCEEFC=CCCD:@A@A@>:>:??B@@B>?@C8@@>>>::>>
@HWI-ST560:155:C574EACXX:3:1101:3129:1977 1:N:0:
CACAGGGAGGACGATGCGGAAAGGAGCTACATTGTCTGCTGGGTTTCAAGAAATCTGACTTTGCTCAGGGAGCCCTGCGTGTCAGTCACTTCCAGCGGTC
+CACAGGGAGGACGATGCGGAAAGGAGCTACATTGTCTGCTGGGTTTCAAGAAATCTGACTTTGCTCAGGGAGCCCTGC
BBCFFFFFHHHHHJJJJJJJJJJJIJJJJJJJJJJJJJJJJJJHIJJJJJJJHHHHHHFFFFFFEEEEEDDDDDDCDDDDDDDDDDDDDDDDDCA?9><@
@HWI-ST560:155:C574EACXX:3:1101:3058:1989 1:N:0:
GCTAGGGAGGACGATGCGGCTAAGAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGCGTGTCAGTCACTTCCAGCGGTCGTATGC
+GCTAGGGAGGACGATGCGGCTAAGAGCTTCCGGGAAACCAAAGTCTTTGGGTTCCGGGGGGAGTATGGTTGC
CCCFFFFFHHHHHGHIJJJIJIJFIIIIIJJIJJJJJIJCHIJHHHHHHGDDFEECDDDDB<BBCCDCDDDDDDDBDCDCCDDDDDDDDCDDBDDB8>9@
@HWI-ST560:155:C574EACXX:3:1101:3093:1995 1:N:0:
AGCAGGGAGGACGATGCGGTTAAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAGAAAAAAAAAAAAGAACAACA
+AGCAGGGAGGACGATGCGGTTAA
B@BFFFFFHHHHFIHIHIJJJJJICEIIIJJJJJJIGIJHIIIGHGEHHCEFB?ACB>@@ACDC?CCC9@##############################
@HWI-ST560:155:C574EACXX:3:1101:3383:1937 1:N:0:
NACAGGGAGGACGATGCGGAAATGATGATGGCTAGAAAAATGAAAGACACAGATAGCGAAGAAGAGATCCGCGAGGCCTTCCGAGTGTTTGACGTGTCAG
+NACAGGGAGGACGATGCGGAAATGATGATGGCTAGAAAAATGAAAGACACAGATAGCGAAGAAGAGATCCGCGAGGCCTTCCGAGTGTTTGAC
#1=DDDFFHHHHHJJJJJJJJJJIJIIJJIJIJIIIIGIJGGIIIGIJIJJIIHHHHFFFDDDCDDDDDDDDDD>@BBDDDDDBB<ACEDDDDB?@CD##
@HWI-ST560:155:C574EACXX:3:1101:3306:1945 1:N:0:
GCTAGGGAGGACGATGCGGTAAGGGTAGTGTTTCCTACTTTATGGAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+GCTAGGGAGGACGATGCGGTAAGGGTAGTGTTTCCTACTTTATGGA
@CCFFFFFHGHHHIIIIJJJJJJIIGHIGGIJJJIIJIJIJFIFHGIIJJJIIIJIJJJHHHHHFFDDDDDDDDCDDDDDDDDDDDDDDDDDD@BDDD9@
@HWI-ST560:155:C574EACXX:3:1101:3255:1945 1:N:0:
GCTAGGGAGGACGATGCGGAAGCGTACAGTACTGTGATAGCTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAAAAAA
+GCTAGGGAGGACGATGCGGAAGCGTACAGTACTGTGATAGCT
CCCFFFFFHHHHHIJJJJJJJJJJHJJJJHIIIJIJI@CCHGCHIJJJIGIIJHGEHGH>BCBDDBDDDDDDDDDDDDDDDDDDDDDD>95555<@@>B5
@HWI-ST560:155:C574EACXX:3:1101:3370:1951 1:N:0:
AACAGGGAGGACGATGCGGAAATGATGATGGCTAGAAAAATGAAAGACACAGATAGCGAAGAAGAGATCCGCGAGGCCTTCCGAGTGTTTGACGTGTCAG
+AACAGGGAGGACGATGCGGAAATGATGATGGCTAGAAAAATGAAAGACACAGATAGCGAAGAAGAGATCCGCGAGGCCTTCCGAGTGTTTGAC
11144222<1?DF>C@GIB?D??BDE9<D@>BFCF9D<;;B<FACC=FE37=@CCA;BED>9@??;?CCCA?6;-9&8>CD(5<@<?+:?9>>@8?C?>3
@HWI-ST560:155:C574EACXX:3:1101:3452:1956 1:N:0:
GCTAGGGAGGACGATGCGGCGATGATGAAAGTGTCACTCGGTGCTAAGCAGGAGACTTCGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG
+GCTAGGGAGGACGATGCGGCGATGATGAAAGTGTCACTCGGTGCTAAGCAGGAGACTTC
@@@FFFFFHHHHHIIIHHIIJIIEHJEIJJJGGGFHIEHEHCHFFFFDEEC?CBDDDDD?CDDBCDDDDDDDDDDDDBD@<BB<CDDDDBBDDDDDDDDC
@HWI-ST560:155:C574EACXX:3:1101:3331:1961 1:N:0:
GTTCCATCAACATCATAGCCAGATGCCCAGAGATTAGAGCGCATGACAAGTAAAGGACGGTTGTCAGCGTCATAAGAGGTTTTACCTCCAAATGAAGAAA
+GTTCCATCAACATCATAGCCAGATGCCCAGAGATTAGAGCGCATGACAAGTAAAGGACGGTTGTCAGCGTCATAAGAGGTTTTACCTCCAAATGAAGAAA
CCCFFBFFFHHGHJIIJIIJJJJJIJIJJJJIJJJJGJIIJJJJJJJJIJGHIIHEHIIIHGHGFFFFFDDDDEDDDDDCBDDDDDDDDCDDDCAC::AC
@HWI-ST560:155:C574EACXX:3:1101:3262:1970 1:N:0:
CACAGGGAGGACGATGCGGTGAGGAGTTTCTGACTATAATGAATAGTTACTATGAACATAGTACTCAAGGCGTGTCAGTCACTTCCAGCGGTCGTATGCC
+CACAGGGAGGACGATGCGGTGAGGAGTTTCTGACTATAATGAATAGTTACTATGAACATAGTACTCAAGGC
B@CFFFFFHHHHHIIIHIIJJJJJIJJIJIJJGJHGIJIIHHIJJJJIJJJJJIJIJIJJIEHHHFHGFFFDDDDDDDDDDDDD>>@CCBDBDDDACDCC
@HWI-ST560:155:C574EACXX:3:1101:3487:1973 1:N:0:
AGCAGGGAGGACGATGCGGATAAGTAGCTTATCAGACTGATGTTGACGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAA
+AGCAGGGAGGACGATGCGGATAAGTAGCTTATCAGACTGATGTTGAC
BBCFFFFFHHHHHJJJHIJJJJJIDGIJJJJJJJJJJJJJIJJJIIJIJJJJJJHHHHHHFFFFFCDDDDDBDCDAB@DDDDCCDDDCDDDDDDD@DBDD
@HWI-ST560:155:C574EACXX:3:1101:3299:1974 1:N:0:
CACAGGGAGGACGATGCGGTCAAGAGACCGTGAAAGCGGGGCCTCACGATCCTTCTGACCTTTTGGGTTTTAAGCAGGAGGTGTCAGAAAAGTTACCACG
+CACAGGGAGGACGATGCGGTCAAGAGACCGTGAAAGCGGGGCCTCACGATCCTTCTGACCTTTTGGGTTTTAAGCAGGAGGTGTCAGAAAAGTTACCACG
BB@DFFFFGHHFHIIJJJJJIIJFHIIIJJIJJJJFDBHIDDDDDDDBDDDDDDDDDEDDDDDDDDDDDDDDC>>>ABDDD?BDDCDDDDDCCC######
@HWI-ST560:155:C574EACXX:3:1101:3449:1974 1:N:0:
CACAGGGAGGACGATGCGGAAGAAGTAGTGTTTCCTACTTTATGGAGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+CACAGGGAGGACGATGCGGAAGAAGTAGTGTTTCCTACTTTATGGA
@@@FFFFFHHHHGIIGGGGBHBFGAGG>9B9DABFDGCGGG9CG<F28CCAC4==DDHEAEE:@4=ABB?AC5>>=;98@CA:>CA>A@C##########
@HWI-ST560:155:C574EACXX:3:1101:1179:1979 1:N:0:
AGCAGGGAGGACGATGCGGGTCAGTCACTTCCAGCGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAACAATATAATAACA
+AGCAGGGAGGACGATGC
@@CFFFFFDHDHFIGIGHIF?DHH?FGIJIIIIIII;FAEGC@EEHHG9EDFFDB'9;8ACCCAB;B8<:>343>483:A####################
@HWI-ST560:155:C574EACXX:3:1101:4452:1943 1:N:0:
CACAGGGAGGACGATGCGGAGGAGAAGACCACATATGTGAAGGCCCCTGGTTGACTGGTTGTGGGCTCAGCTGACCAGCTGGGCTTGCCTGCTGCAGGCG
+
BBBFFFFFHHHHHJJJJJJJIJGGIJFBGH@GGHIGGDHIICHIJJHHHGFFFCDEEDCED@@@DDDDDDCCACDD@BBCCDDBCDDDDDDDDCA99<9@
`
