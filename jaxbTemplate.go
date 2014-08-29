package main

type JaxbInfo struct {
	Name                   string
	Root                   bool
	PackageName, ClassName string
	Attributes             []*JaxbAttribute
	Fields                 []*JaxbField
	HasValue               bool
}

type JaxbAttribute struct {
	Name      string
	NameUpper string
	NameLower string
	NameSpace string
}
type JaxbField struct {
	TypeName  string
	Name      string
	NameUpper string
	NameLower string
	NameSpace string
	Repeats   bool
}

func (jb *JaxbInfo) init() {
	jb.Attributes = make([]*JaxbAttribute, 0)
	jb.Fields = make([]*JaxbField, 0)
}

const jaxbTemplate = `
// Generated by chidley https://github.com/gnewton/chidley
package {{.PackageName}};

import java.util.ArrayList;
import javax.xml.bind.annotation.*;

@XmlAccessorType(XmlAccessType.FIELD)
@XmlRootElement(name="{{.Name}}")
public class {{.ClassName}} {
{{if .Attributes}}
    // Attributes{{end}}
{{range .Attributes}}
{{if .NameSpace}}    
@XmlAttribute(namespace = "{{.NameSpace}}"){{else}}    @XmlAttribute{{end}}
    public String {{.NameLower}};{{end}}
{{if .Fields}}
    // Fields{{end}}{{range .Fields}}    
    @XmlElement
    {{if .Repeats}}public ArrayList<{{.TypeName}}> {{.NameLower}}{{else}}public {{.TypeName}} {{.NameLower}}{{end}};
{{end}}
{{if .HasValue}}
    // Value
    @XmlValue
    public String text;{{end}}
}
`

const jaxbMainTemplate = `
// Generated by chidley https://github.com/gnewton/chidley
package ca.newtong.chidley;
 
import java.io.File;
import javax.xml.bind.JAXBContext;
import javax.xml.bind.JAXBException;
import javax.xml.bind.Unmarshaller;
 
public class Main {
	public static void main(String[] args) {
	 try {
		File file = new File("/home/newtong/work/chidley/xml/Fantasia_con_imitazione_BWV563.xml");
		JAXBContext jaxbContext = JAXBContext.newInstance(Score_partwise.class);
 
		Unmarshaller jaxbUnmarshaller = jaxbContext.createUnmarshaller();
		Score_partwise score = (Score_partwise) jaxbUnmarshaller.unmarshal(file);
		System.out.println(score);
	  } catch (JAXBException e) {
		e.printStackTrace();
	  }
	}
}
`
