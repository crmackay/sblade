package main

import (
    "fmt"
    bio "github.com/crmackay/gobioingfo"
    
    
struct TrimmedFASTQRead {
    bio.FASTQRead
    CutSequence string // place to store the cut string
}
    
func Worker(input chan bio.FASTQRead, ouput chan bio.FASTQRead) {

    var doneRead TrimmedFASTQRead
    
    for read := range input {
        
        loop:
            for {
             alignment := align(read, linker)
             
             hasContaminant := BayesTest(alignmment)
             
             if hasContaminant == true {
                //
             }
            
            }
        
       
        
        
        // align linker
        // Bayesian test
        // align linker
        // Bayesian test
        // cut
        // put both cut and full length into output queue
        
        output <- cutRead
        output <- fragment
        
    }

}