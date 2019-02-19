# neke
Generates interfaces and structs for Go and TypeScript.

## Getting started
```
go get -v github.com/32leaves/neke
neke help
```

You can play also with Neke in Gitpod:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io#github.com/32leaves/neke)


## Example

```
interface VisitorCountService {
    func countVisitor (common:CountVisitorRequest) returns (comon:CountVisitorResponse)
    func getVisitorCount (common:GetVisitorCountRequest) returns (common:CountVisitorResponse)
}

struct CountVisitorRequest {
    optional name   common:string
    optional weight go:uint16     ts:number
    required time   go:time.Time  ts:Date
    required kind   common:VisitorKind
}

struct CountVisitorResponse {
    required count go:uint64 ts:number
}

struct GetVisitorCountRequest { }

enum VisitorKind {
    FirstTimeVisitor
    RepeatedVisitor
    LongTimeVisitor
}
```

turns to Go code using `neke generate -l go examples/readme.neke`:
```Go
type VisitorCountService interface {
        CountVisitor(req *CountVisitorRequest) (*CountVisitorResponse, error)
        GetVisitorCount(req *GetVisitorCountRequest) (*CountVisitorResponse, error)
}

type CountVisitorRequest struct {
        Name   string      `json:"name,omitempty"`
        Weight uint16      `json:"weight,omitempty"`
        Time   string      `json:"time"`
        Kind   VisitorKind `json:"kind"`
}

type CountVisitorResponse struct {
        Count uint64 `json:"count"`
}

type GetVisitorCountRequest struct {
}

type VisitorKind string

const (
        VisitorKind_FirstTimeVisitor VisitorKind = "firstTimeVisitor"
        VisitorKind_RepeatedVisitor  VisitorKind = "repeatedVisitor"
        VisitorKind_LongTimeVisitor  VisitorKind = "longTimeVisitor"
)
```

and TypeScript using `neke generate -l ts examples/readme.neke`:
```TypeScript
export interface VisitorCountService {
    CountVisitor(req CountVisitorRequest): CountVisitorResponse;
    GetVisitorCount(req GetVisitorCountRequest): CountVisitorResponse;
}

export interface CountVisitorRequest {
    name?: string;
    weight?: number;
    time: Date;
    kind: VisitorKind;
}

export interface CountVisitorResponse {
    count: number;
}

export interface GetVisitorCountRequest {
}

export enum VisitorKind {
    FirstTimeVisitor = "firstTimeVisitor";
    RepeatedVisitor = "repeatedVisitor";
    LongTimeVisitor = "longTimeVisitor";
}
```

## Contribution
Issues and PRs are always welcome ❤️.

The free online IDE Gitpod makes it super easy to help with this project:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io#github.com/32leaves/neke)