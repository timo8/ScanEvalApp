package common

const (
	// Specifies the permission mode used for generated files (read/write for owner, read-only for group and others).
	FILE_PERMISSION = 0644

	// Path to the main LaTeX template used for PDF generation
	TEMPLATE_PATH = "./assets/latex/main.tex"

	// Directory where generated PDF files are stored
	OUTPUT_PDF_PATH = "./assets/tmp"

	// Temporary directory used during PDF generation for LaTeX compiling purposes
	TEMPORARY_PDF_PATH = "./assets/tmp"

	// Export dir path where to store all kind of PDFs
	EXPORT_DIR = "./assets/tmp/"
)

// TODO - prerobit nejako, aby sa vracali normalne spravy, ked nastane error
const (
	StatusSuccess        = 0
	StatusFileNotFound   = 1
	StatusInvalidFormat  = 2
	StatusProcessingFail = 3
)
