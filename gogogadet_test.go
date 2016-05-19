package gogogadjet

import (
	//"fmt"
	"testing"
	"bytes"
	"strings"
)

func Test_Init(t *testing.T) {
	MakeDir("Test")
	WriteStringToFile("Test/Test_FileExists.txt", "Test Data")
	MakeDir("Test/FolderExists_Test")
	MakeDir("Test/Test_FileDelete")
	WriteStringToFile("Test/Test_FolderExistsFile.txt", "Test Data")
	WriteStringToFile("Test/Test_ReadStringOfFile.txt", "Testing \r\nSome Text\r\nTo Lines\r\nDone")
	WriteStringToFile("Test/Test_ReadLinesOfFile.txt", "Testing \r\nSome Text\r\nTo Lines\r\nDone")
	WriteStringToFile("Test/Test_ReadByteSliceOfFile.txt", "Testing \r\nSome Text\r\nTo Lines\r\nDone")
}

func Test_NewGuid(t *testing.T) {
	guid := NewGuid()
	if (guid.String() == "") {
		t.Error("Guid Is Blank")
	}
}

func Test_StringToByteSlice(t *testing.T) {
	b := StringToByteSlice("Testing \r\n")
	
	if (bytes.Equal(b, []byte{84,101,115,116,105,110,103,32,13,10}) == false) {
		t.Error("Bytes Dont Match")
	}
}

func Test_ByteSliceToString(t *testing.T) {
	s := ByteSliceToString([]byte{84,101,115,116,105,110,103,32,13,10})
	
	if (strings.Compare(s, "Testing \r\n") != 0) {
		t.Error("Strings Dont Match")
	}
}

func Test_ReadLinesOfString(t *testing.T) {
	s := "Testing \r\nSome Text\nTo Lines\r\nDone"
	
	lines, err := ReadLinesOfString(s)
	if (err != nil) {
		t.Error("Error: " + err.Error())
	} else if (len(lines) != 4) {
		t.Error("Lines != 4: ", len(lines))
	} else {
		if (strings.Compare(lines[0], "Testing ") != 0) {
			t.Error("String Dont Match: '" + lines[0] + "'")
		}
		if (strings.Compare(lines[1], "Some Text") != 0) {
			t.Error("String Dont Match: '" + lines[1] + "'")
		}
		if (strings.Compare(lines[2], "To Lines") != 0) {
			t.Error("String Dont Match: '" + lines[2] + "'")
		}
		if (strings.Compare(lines[3], "Done") != 0) {
			t.Error("String Dont Match: '" + lines[3] + "'")
		}
	}
}

func Test_ReadLinesOfByteSlice(t *testing.T) {
	s := "Testing \r\nSome Text\nTo Lines\r\nDone"
	b := StringToByteSlice(s)
	
	lines, err := ReadLinesOfByteSlice(b)
	if (err != nil) {
		t.Error("Error: " + err.Error())
	} else if (len(lines) != 4) {
		t.Error("Lines != 4: ", len(lines))
	} else {
		if (strings.Compare(lines[0], "Testing ") != 0) {
			t.Error("String Dont Match: '" + lines[0] + "'")
		}
		if (strings.Compare(lines[1], "Some Text") != 0) {
			t.Error("String Dont Match: '" + lines[1] + "'")
		}
		if (strings.Compare(lines[2], "To Lines") != 0) {
			t.Error("String Dont Match: '" + lines[2] + "'")
		}
		if (strings.Compare(lines[3], "Done") != 0) {
			t.Error("String Dont Match: '" + lines[3] + "'")
		}
	}
}

func Test_FileExists(t *testing.T) {
	b, err := FileExists("Test/Test_FileExists.txt")
	if (err != nil) {
		t.Error("Error 1: " + err.Error())
	} else {
		if (b == false) {
			t.Error("Error 2")
		}
	}
	
	b, err = FileExists("Test/Test_FileExistsNot.txt")
	if (err != nil) {
		t.Error("Error 3: " + err.Error())
	} else {
		if (b == true) {
			t.Error("Error 4")
		}
	}
	
	b, err = FileExists("Test")
	if (err	 == nil) {
		t.Error("Error 5")
	}
}

func Test_MakeDir(t *testing.T) {
	err := MakeDir("Test/Test_MakeDir")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		b, err := FolderExists("Test/Test_MakeDir")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			if (b == false) {
				t.Error("Error 3")
			}
		}
	}
	
	err = MakeDir("Test/Test_MakeDir")
	if (err == nil) {
		t.Error("Error 1")
	} else {
		b,err := FolderExists("Test/Test_MakeDir")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			if (b == false) {
				t.Error("Error 3")
			}
		}
	}
}

func Test_FolderExists(t *testing.T) {
	b,err := FolderExists("Test/FolderExists_Test")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		if (b == false) {
			t.Error("Error 2")
		}
	}
	
	b, err = FolderExists("Test/FolderNotExists_Test")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		if (b == true) {
			t.Error("Error 2")
		}
	}
	
	b, err = FolderExists("Test/Test_FolderExistsFile.txt")
	if (err == nil) {
		t.Error("Error 1")
	}
}

func Test_FileDelete(t *testing.T) {
	err := WriteStringToFile("Test/Test_FileDelete/Test.txt", "Test Data")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		err := FileDelete("Test/Test_FileDelete/Test.txt")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			b,err := FileExists("Test/Test_FileDelete/Test.txt")
			if (err != nil) {
				t.Error("Error 3:" + err.Error())
			} else {
				if (b == true) {
					t.Error("Error 4")
				}
			}
		}
	}
	
	//this tests if i get error when tryng to delete file that does not exist
	err = FileDelete("Test/Test_FileDelete/TestNoFile.txt")
	if (err == nil) {
		t.Error("Error 1")
	}
}

func Test_FolderDelete(t *testing.T) {
	err := MakeDir("Test/Test_FolderDelete")	
	if (err	!= nil) {
		t.Error("Error 1:" + err.Error())
	}
	
	//this is testing if get error when delete folder with data in it and false sent
	err = WriteStringToFile("Test/Test_FolderDelete/Test.txt", "some data")
	if (err != nil) {
		t.Error("Error 2:" + err.Error())
	} else {
		err = FolderDelete("Test/Test_FolderDelete", false)
		if (err == nil) {
			t.Error("Error 3")
		} 
	}
	
	//test is testing if i can delete the same folder with true set
	err = FolderDelete("Test/Test_FolderDelete", true)
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		b, err := FolderExists("Test/Test_FolderDelete")
		if (err != nil)	{
			t.Error("Error 3:" + err.Error())
		} else {
			if (b == true) {
				t.Error("Error 4")
			}
		}
	}
}

func Test_WriteStringToFile(t *testing.T) { 
	s := "Testing \r\nSome Text\nTo File"
	err := WriteStringToFile("Test/Test_WriteStringToFile.txt", s)
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		s2, err := ReadStringOfFile("Test/Test_WriteStringToFile.txt")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			if (strings.Compare(s, s2) != 0) {
				t.Error("Error 3")
			}
		}
	}
}

func Test_WriteAllLinesToFile(t *testing.T) { 
	lines := []string{"Testing ","Some Text","To Lines","Done"}
	
	err := WriteAllLinesToFile("Test/Test_WriteAllLinesToFile.txt", lines)
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		rlines,err := ReadLinesOfFile("Test/Test_WriteAllLinesToFile.txt")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			if (len(lines) != 4) {
				t.Error("Error 4")
			} else {
				if (strings.Compare(lines[0], rlines[0]) != 0) {
					t.Error("Error 5: '" + rlines[0] + "'")
				}
				if (strings.Compare(lines[1], rlines[1]) != 0) {
					t.Error("Error 6: '" + rlines[1] + "'")
				}
				if (strings.Compare(lines[2], rlines[2]) != 0) {
					t.Error("Error 7: '" + rlines[2] + "'")
				}
				if (strings.Compare(lines[3], rlines[3]) != 0) {
					t.Error("Error 8: '" + rlines[3] + "'")
				}
			}
		}
	}
}

func Test_WriteByteSliceToFile(t *testing.T) { 
	b := StringToByteSlice("Some Test Text")
	
	err := WriteByteSliceToFile("Test/Test_WriteByteSliceToFile.txt", b)
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		bread,err := ReadLinesOfFile("Test/Test_WriteByteSliceToFile.txt")
		if (err != nil) {
			t.Error("Error 2:" + err.Error())
		} else {
			if (strings.Compare(bread[0], "Some Test Text") != 0) {
				t.Error("Error 3")
			}
		}
	}
}

func Test_ReadLinesOfFile(t *testing.T) { 
	lines, err := ReadLinesOfFile("Test/Test_ReadLinesOfFile.txt")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else if (len(lines) != 4) {
		t.Error("Error 2")
	} else {
		if (strings.Compare(lines[0], "Testing ") != 0) {
			t.Error("Error 3: '" + lines[0] + "'")
		}
		if (strings.Compare(lines[1], "Some Text") != 0) {
			t.Error("Error 3: '" + lines[1] + "'")
		}
		if (strings.Compare(lines[2], "To Lines") != 0) {
			t.Error("Error 3: '" + lines[2] + "'")
		}
		if (strings.Compare(lines[3], "Done") != 0) {
			t.Error("Error 3: '" + lines[3] + "'")
		}
	}
}

func Test_ReadByteSliceOfFile(t *testing.T) { 
	bfile, err := ReadByteSliceOfFile("Test/Test_ReadByteSliceOfFile.txt")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		s := "Testing \r\nSome Text\r\nTo Lines\r\nDone"
		b := StringToByteSlice(s)
		
		if (bytes.Equal(bfile, b) == false) {
			t.Error("Error 2")
		}
	}
}

func Test_ReadStringOfFile(t *testing.T) { 
	bfile, err := ReadByteSliceOfFile("Test/Test_ReadStringOfFile.txt")
	if (err != nil) {
		t.Error("Error 1:" + err.Error())
	} else {
		s := "Testing \r\nSome Text\r\nTo Lines\r\nDone"
		
		if (bytes.Equal(bfile, StringToByteSlice(s)) != true) {
			t.Error("Error 2")
		}
	}
}

func Test_Cleanup(t *testing.T) {
	FolderDelete("Test", true)
}
