package response_types

// apisdkCreateThingResponseBool illustrates handling for primitive types (i.e., *bool, *float, *int, *string).
type apisdkCreateThingResponseBool struct {
	Bool *bool
}

// apisdkCreateThingResponseList illustrates handling for list type.
type apisdkCreateThingResponseList struct {
	List []*string
}

// apisdkCreateThingResponseList illustrates handling for map type.
type apisdkCreateThingResponseMap struct {
	Map map[string]*bool
}

// apisdkCreateThingResponseSet illustrates handling for set type.
type apisdkCreateThingResponseSet struct {
	List []*string
}

// apisdkCreateThingResponseObject illustrates handling for object type.
type apisdkCreateThingResponseObject struct {
	Object *apisdkObject
}

// apisdkObject
type apisdkObject struct {
	Property *string
}
