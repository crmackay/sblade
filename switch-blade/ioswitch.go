package main

import (
        "fmt"
        bio "github.com/crmackay/gobioinfo"
        "path/filepath"
)


// this function uses the `select` method to keep IO to just one goroutine
// this combined with our CPU accounting allows for fast loading and offloading
// of data to disk, while ensuring that that maximum CPU utilization is happening


func (r *TrimmedFASTQRead) Seperate (t bio.FASTQRead, f bio.FASTQRead) {

    trimPosition := len(r.Sequence) - len(r.CutSequence)

    t = bio.FASTQRead{Name: r.Name, Sequence: r.CutSequence , Misc: Misc, Quality:r.Quality[0:trimPosition]}
    
    f = bio.FASTQRead{Name: r.Name, Sequence: r.Sequence[trimPosition:] , Misc: Misc, Quality:r.Quality[trimPosition:]}

}

func IOSwitch(inFile string, outFile string, input chan bio.FASTQRead, output chan bio.FASTQRead) {
    
    reader := bio.NewFASTQScanner(inFile)
    defer reader.Close()
    
    cutWriter := bio.NewFASTQWriter(outFile)
    defer cutWriter.Close()
    
    piecesWriter := bio.NewFASTQWriter(filepath.Dir(outFile)+"/removed_pieces.fastq")
    defer piecesWriter.Close()
    
    var newRead, doneRead bio.FASTQRead
    
    // TODO: how do you detect when the file is done?...
    
    loop: 
        for {
            
            newRead = reader.NextRead()
            
            if newRead.Sequence == "" {
                
                for i := 0; i < len(output); i++ {
                    writer.Write(<- output)
                }
                
                break loop
            }
            
            select{
            
            case input <- newRead:
                //input is not full, great fill it up!
                
            default:
                //once the input is full, drain the output
                for i := 0; i < len(output); i++ {
                    doneRead <- output
                    
                    cutRead, fragment := Seperate(doneRead)
                    
                        cutWriter.Write(cutRead)
                        piecesWriter.Write(fragment)
                    }
                    
                }
            }
        }
}




