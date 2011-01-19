package netlink

import "encoding/binary"
import "os"

type AttributeFinder interface {
  GetAttribute(AttributeType)(Attribute, os.Error)
}

func GetAttributeUint32(af AttributeFinder, at AttributeType)(out uint32, err os.Error){
  attr, err := af.GetAttribute(at)
  if err == nil {
    body := attr.Body()
    if len(body) != 4 {
      err = os.NewError("Attribute wrong size for Uint32")
    } else {
      out = binary.LittleEndian.Uint32(body)
    }
  }
  return
}

func GetAttributeString(af AttributeFinder, at AttributeType)(out string, err os.Error){
  attr, err := af.GetAttribute(at)
  if err == nil {
    out = string(attr.Body())
  }
  return
}

// Same as GetAttributeString, but expects the string to be NULL terminated.
func GetAttributeCString(af AttributeFinder, at AttributeType)(out string, err os.Error){
  attr, err := af.GetAttribute(at)
  if err == nil {
    outbody := attr.Body()
    if len(outbody) == 0 {
      err = os.NewError("Invalid body")
    } else {
      if outbody[len(outbody) -1 ] != 0 {
        err = os.NewError("Expected NULL-terminated string")
      } else {
        out = string(outbody[0: len(outbody) - 1])
      }
    }
  }
  return
}