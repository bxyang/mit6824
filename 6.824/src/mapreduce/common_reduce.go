package mapreduce

import (
    "encoding/json"
    "io/ioutil"
    "strings"
    "sort"
    "os"
)


func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	//
	// doReduce manages one reduce task: it should read the intermediate
	// files for the task, sort the intermediate key/value pairs by key,
	// call the user-defined reduce function (reduceF) for each key, and
	// write reduceF's output to disk.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTask) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Your code here (Part I).
	//

    var kv KeyValue     
    
    data := make(map[string][]string)
    keys := make(sort.StringSlice, 0)

    for i := 0; i < nMap; i++ {
        mFileName := reduceName(jobName, i, reduceTask)        
        content, _ := ioutil.ReadFile(mFileName)
        arrs := strings.Split(string(content), "\n")
        for j, item := range arrs {
            if j == len(arrs) - 1 {
                break
            }
            dec := json.NewDecoder(strings.NewReader(item))
            dec.Decode(&kv)

            _, ok := data[kv.Key]
            if !ok {
                data[kv.Key] = make([]string, 0)
                keys = append(keys, kv.Key)
            }
            data[kv.Key] = append(data[kv.Key], kv.Value)
        }
    }
    sort.Sort(keys)
    
    ofile, _ := os.Create(outFile)
    enc := json.NewEncoder(ofile)
    
    for _, k := range(keys) {
        v := reduceF(k, data[k])
        enc.Encode(KeyValue{k, v})
    }
    
    ofile.Close()
}
