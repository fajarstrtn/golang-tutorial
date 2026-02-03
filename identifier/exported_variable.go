package identifier

/*
 * The identifier which is allowed to access it
 * from another package is known as the exported identifier.
 * The exported identifiers are those identifiers which obey the following condition:
 * 1. The first character of the exported identifier's name
 * should be in the Unicode upper case letter.
 * 2. The identifier should be declared in the package block or be a variable,
 * function, type, or method name within that package.
 *
 * The uniqueness of the identifiers means the identifier is unique from
 * the other set of the identifiers available in your program,
 * or in the package and they are not exported.
 *
 * The ExportedVariable is recognized and accessible in its package. */
var ExportedVariable string = "This is an exported variable"
