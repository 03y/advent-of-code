package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type directory struct {
    Name        string
    Directories []*directory
    Files       []file
    Parent      *directory
    Size        int64
}

type file struct {
    Name string
    Size int64
}

const diskSpace, spaceNeeded, maxDirSize = 70_000_000, 30_000_000, 100_000

func main() {
    input, _ := readLines("input.txt")
    root := buildTree(input)
    unusedSpace := diskSpace - root.Size
    spaceToFreeUp := spaceNeeded - unusedSpace

    filesystemSize, spaceToDelete := answer(&root, spaceToFreeUp, root.Size)
    fmt.Println("Part 1 Answer:", filesystemSize)
    fmt.Println("Part 2 Answer:", spaceToDelete)
}

func answer(dir *directory, spaceToFreeUp int64, spaceToDelete int64) (int64, int64) {
    var filesytemSize int64
    var size = spaceToDelete // size of the smallest directory that can be deleted to give enough space

    if dir.Size <= maxDirSize {
        filesytemSize += dir.Size
    }

    if dir.Size >= spaceToFreeUp && dir.Size < spaceToDelete {
        size = dir.Size
    }

    for _, subdirs := range dir.Directories {
        newFsSize, newBestCandidate := answer(subdirs, spaceToFreeUp, size)
        filesytemSize += newFsSize
        size = newBestCandidate
    }

    return filesytemSize, size
}

func buildTree(commands []string) directory {
    var root directory
    root.Name = "/"
    var currentDir = &root

    for _, instruction := range commands {
        command := strings.Split(instruction, " ")
        if strings.HasPrefix(instruction, "$") {
            if command[1] == "cd" {
                if command[2] == ".." {
                    currentDir = currentDir.Parent
                } else {
                    for _, dir := range currentDir.Directories {
                        if dir.Name == command[2] {
                            currentDir = dir
                        }
                    }
                }
            }
        } else {
            if command[0] == "dir" {
                dir := directory{Name: command[1], Parent: currentDir}
                currentDir.Directories = append(currentDir.Directories, &dir)
            } else {
                size, _ := strconv.ParseInt(command[0], 10, 64)
                newFile := file{Name: command[1], Size: size}
                currentDir.Files = append(currentDir.Files, newFile)
                enlargeFilesystem(size, currentDir)
            }
        }
    }
    return root
}

func enlargeFilesystem(size int64, dir *directory) {
    dir.Size += size
    if dir.Parent != nil {
        enlargeFilesystem(size, dir.Parent)
    }
}

func readLines(path string) ([]string, error) {
    file, _ := os.Open(path)
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
