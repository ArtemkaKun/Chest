package main

import (
	"fmt"
	"log"
	"time"
)

func CheckOld(backup_folder string, max_files int) {
	service, err := getService()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := service.Files.List().Q(fmt.Sprintf("'%v' in parents", backup_folder)).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		if len(r.Files) > max_files {
			for range r.Files {
				for n, j := range r.Files {
					file_date, err := time.Parse("2006-01-02", j.Name[0:10])
					if err != nil {
						log.Fatalf("%v", err)
					}
					if n+1 <= len(r.Files)-1 {
						next_file_date, _ := time.Parse("2006-01-02", r.Files[n+1].Name[0:10])

						if file_date.After(next_file_date) {
							buffer := r.Files[n]
							r.Files[n] = r.Files[n+1]
							r.Files[n+1] = buffer
						}
					}
				}
			}

			for _, k := range r.Files[0 : len(r.Files)-max_files] {
				service.Files.Delete(k.Id).Do()
				fmt.Printf("%v (%vs) was deleted\n", k.Name, k.Id)
			}
		}
	}
}
