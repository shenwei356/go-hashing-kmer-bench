#!/bin/sh

csvtk space2tab result.txt \
    | csvtk sort -Ht -k 3:n \
    | csvtk mutate2 -Ht -e '$3 + " " + $4' \
    | csvtk cut -Ht -f 1,2,5 \
    | csvtk -Ht pretty -s '        '
