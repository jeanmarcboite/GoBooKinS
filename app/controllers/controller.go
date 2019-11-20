package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"

	"github.com/revel/revel"
)

// Controller -- application controller
type Controller struct {
	*revel.Controller
}

// SprintHTML  -- print data in HTML
func (c Controller) SprintHTML(x interface{}) template.HTML {
	xm, err := json.Marshal(x)
	if err != nil {
		c.Log.Errorf("Marshal error: err")
		return template.HTML(err.Error())
	}
	var xmu interface{}
	err = json.Unmarshal(xm, &xmu)
	if err != nil {
		c.Log.Errorf("Unmarshal error: %v", err)
		return template.HTML(err.Error())
	}

	return template.HTML(styleSheet+c.toHTML(xmu, 1))
}

func (c Controller) toHTML(x interface{}, id int) string {
	c.Log.Debugf("toHTML %v", reflect.TypeOf(x))
	switch v := x.(type) {
	case map[string]interface{}:
		return c.mapToHTML(v, id)
	case []interface{}:
		return c.arrayToHTML(v, id)
	}

	return fmt.Sprintf("%v", x)
}

/*
   <li><input type="checkbox" id="c1" />
       <i class="fa fa-angle-double-right"></i>
       <i class="fa fa-angle-double-down"></i>
       <label for="c1">Dossier A</label>
       <ul>
           <li>Sous dossier A1</li>
           <li>Sous dossier A2</li>
           <li>Sous dossier A3</li>
       </ul>
   </li>
*/

func (c Controller) mapToHTML(m map[string]interface{}, id int) string {
	checkbox := `
     <li><input type='checkbox' id='__c%v' />
        <i class='fa fa-angle-double-right'></i>
        <i class='fa fa-angle-double-down'></i>
        <label for='__c%v'>%v</label>
        %v
    </li>
    `
	value := `<li>"%v": "%v"</li>`

	bufferString := bytes.NewBufferString("<ul>")
	for k, v := range m {
		id++
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			fmt.Fprintf(bufferString, checkbox, id, id, k, c.toHTML(v, id))
		default:
			fmt.Fprintf(bufferString, value, k, c.toHTML(v, id))
		}
	}
	bufferString.WriteString("</ul>")
	return bufferString.String()
}

func (c Controller) arrayToHTML(a []interface{}, id int) string {
	format := `<li>%v</li>`

	bufferString := bytes.NewBufferString("<ul>")
	for _, v := range a {
		id++
		fmt.Fprintf(bufferString, format, c.toHTML(v, id))
	}
	bufferString.WriteString("</ul>")
	return bufferString.String()
}

const styleSheet = `
<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.11.2/css/all.min.css" rel="stylesheet">
<style>
    /* https://makina-corpus.com/blog/metier/2014/construire-un-tree-view-en-css-pur  */
    /* fonctionnel */
    
    input {
        display: none;
    }
    
    input~ul {
        display: none;
    }
    
    input:checked~ul {
        display: block;
    }
    
    input~.fa-angle-double-down {
        display: none;
    }
    
    input:checked~.fa-angle-double-right {
        display: none;
    }
    
    input:checked~.fa-angle-double-down {
        display: inline;
    }
    /* habillage */
    
    li {
        display: block;
        font-family: 'Arial';
        font-size: 15px;
        padding: 0.2em;
        border: 1px solid transparent;
    }
    
    li:hover {
        border: 1px solid grey;
        border-radius: 3px;
        background-color: lightgrey;
    }
</style>
`

const list = `

<ul>
    <li><input type="checkbox" id="c1" />
        <i class="fa fa-angle-double-right"></i>
        <i class="fa fa-angle-double-down"></i>
        <label for="c1">Dossier A</label>
        <ul>
            <li>Sous dossier A1</li>
            <li>Sous dossier A2</li>
            <li>Sous dossier A3</li>
        </ul>
    </li>
    <li><input type="checkbox" id="c2" />
        <i class="fa fa-angle-double-right"></i>
        <i class="fa fa-angle-double-down"></i>
        <label for="c2">Dossier B</label>
        <ul>
            <li>Sous dossier B1</li>
            <li><input type="checkbox" id="c3" />
                <i class="fa fa-angle-double-right"></i>
                <i class="fa fa-angle-double-down"></i>
                <label for="c3">Sous dossier B2</label>
                <ul>
                    <li>Sous-sous dossier B21</li>
                    <li><input type="checkbox" id="c4" />
                        <i class="fa fa-angle-double-right"></i>
                        <i class="fa fa-angle-double-down"></i>
                        <label for="c4">Sous-sous dossier B22</label>
                        <ul>
                            <li>Sous-sous-sous dossier B221</li>
                            <li>Sous-sous-sous dossier B222</li>
                        </ul>
                    </li>
                    <li>Sous-sous dossier B23</li>
                </ul>
            </li>
        </ul>
    </li>
</ul>
`
