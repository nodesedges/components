package pluginmeta

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
 * Version
 * Defines a version according semantic versioning
 * example: 1.0.2-beta+b2.2393220.21-20240116
 */
type Version struct {
	Minor   int    `json:"minor"`
	Major   int    `json:"major"`
	Patch   int    `json:"patch"`
	Release string `json:"release"`
	Build   string `json:"build"`
	Date    string `json:"date"`
}

/**
 * Author
 * Defines an author
 * example: Stefan S., XYZ Comp.
 */
type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Company string `json:"company"`
}

/**
 * License
 * License of the software product
 * example:
 * GNU General Public License, version 1
 *
 * Copyright (C) 1989 Free Software Fou...
 *
 */
type License struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

/**
 * Parameter
 * An parameter has a name and an type.
 * You can choose all JSON Types as type.
 *
 */
type Parameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

/**
 * Function
 * A function definition.
 * Functions have a name, a comment to tell you how to use the function and
 * may have input- and output- parameters.
 */
type Function struct {
	Name      string      `json:"name"`
	Comment   string      `json:"comment"`
	ParamsIn  []Parameter `json:"paramsin"`
	ParamsOut []Parameter `json:"paramsout"`
}

/**
 * Requirement
 * An requirement is used by the plugin to inform the parent application
 * about needed parent API objects and API version in one key.
 *
 * example: "key": "api/v1/management/pushmessage", "access": {"read", "write"}
 * that means the plugin requires access to the "pushmessage" object of the parent application and
 * it likes to use the api "v1"
 */
type Requirement struct {
	ApiKey string `json:"apikey"`
	// Access []string `json:"access"`  // todo: check if needed
}

type CustomType struct {
	Name    string      `json:"name"`
	Comment string      `json:"comment"`
	Vars    []Parameter `json:"vars"`
	// Functions []Function  `json:"functions"` // todo avoid this because of extra overhead, keep apis simple C-like
}

/*
 * Configuration
 * A configuration is a set of properties that could be get and set to configure the use case.
 * A configuration therefore has a value field, a type for the value and a comment, it also indicates if its read only.
 **/
type Configuration struct {
	Name     string `json:"name"`
	Value    any    `json:"value"`
	Comment  string `json:"comment"`
	Writable bool   `json:"writeable"`
}

/**
 * UseCase
 * Each plugin is made for at least one specific usecase.
 * Usecases may have a name, a description, a config and a list of functions which makes the plugin usable by the parent application.
 * Usecases can also have requirements to work correctly in the context of their parent application.
 *
 * !!! Usecases are the most essential part to get loaded and used by the parent application in the right manner !!!
 *
 * We take this approach of en embedded API defintion instead of OpenApi or xtp because of its simplicity.
 * All supported programming languages have native support for JSON so developers can simply write their plugin specification directy in
 * the language of their choise and generate JSON out of it, no need for extra tools or 3-party generators.
 *
 */
type UseCase struct {
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Config       []Configuration `json:"config"`
	Functions    []Function      `json:"functions"`
	Requirements []Requirement   `json:"requirements"`
	CustomTypes  []CustomType    `json:"customtype"`
}

type Description struct {
	// general plugins information
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Version     Version `json:"version"`
	Website     string  `json:"website"`
	Author      Author  `json:"author"`
	License     License `json:"license"`
	// interface definitions
	UseCases []UseCase `json:"usecases"`
}

/**
 * GetDescription
 * must be implemented by the plugin author
 * returns the Plugin Description as JSON string
 */
func GetDescription() string {
	desc := Description{
		Name: "TestPlugin",
		Description: `This is a test plugin for educational use.
feel free to modify and try out everything you like.
					  
For example you can add info from your Readme.md out of 
your project github repository here.
					  
Plugin description will be available in parent application
and is also used for implementing plugin interfaces.
`,
		Version: Version{Minor: 0, Major: 0, Patch: 1, Release: "alpha", Build: "", Date: time.Now().String()},
		Website: "myprojectwebsite.org",
		Author:  Author{Name: "Stefan", Surname: "S.", Company: ""},
		License: License{Name: "The MIT License", Text: `Copyright 2025 Stefan S.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
`},
		// definition of the plugins usecases
		UseCases: []UseCase{{Name: "TestUseCase",
			Description:  "With this usecase you can test the plugin",
			Requirements: nil,
			Functions: []Function{
				{Name: "test",
					Comment:   "call this function, returns the string \"test\"",
					ParamsIn:  nil,
					ParamsOut: []Parameter{{Name: "returnValue", Type: "String"}}},
			},
			Config: []Configuration{
				{
					Name: "Endpoint", Value: "https://nodesedges.com", Comment: "Example default Endpoint for Plugin if needed", Writable: true,
				},
			},
		},
		},
	}

	b, err := json.Marshal(desc)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(b)

}
