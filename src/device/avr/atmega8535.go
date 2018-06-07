// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega8535.atdf, see http://packs.download.atmel.com/

// +build avr,atmega8535

// Device information for the ATmega8535.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega8535"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Reset, Power-on Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt 0
	IRQ_INT1         = 2  // External Interrupt 1
	IRQ_TIMER2_COMP  = 3  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 4  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 5  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 6  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 7  // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 8  // Timer/Counter1 Overflow
	IRQ_TIMER0_OVF   = 9  // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 10 // SPI Serial Transfer Complete
	IRQ_USART_RX     = 11 // USART, RX Complete
	IRQ_USART_UDRE   = 12 // USART Data Register Empty
	IRQ_USART_TX     = 13 // USART, TX Complete
	IRQ_ADC          = 14 // ADC Conversion Complete
	IRQ_EE_RDY       = 15 // EEPROM Ready
	IRQ_ANA_COMP     = 16 // Analog Comparator
	IRQ_TWI          = 17 // Two-wire Serial Interface
	IRQ_INT2         = 18 // External Interrupt Request 2
	IRQ_TIMER0_COMP  = 19 // TimerCounter0 Compare Match
	IRQ_SPM_RDY      = 20 // Store Program Memory Read
	IRQ_max          = 20 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		HIGH __reg
		LOW  __reg
	}{
		HIGH: 0x1,
		LOW:  0x0,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX  __reg
		ADCSRA __reg
		ADCL   __reg
		ADCH   __reg
	}{
		ADMUX:  0x27, // The ADC multiplexer Selection Register
		ADCSRA: 0x26, // The ADC Control and Status register
		ADCL:   0x24, // ADC Data Register  Bytes
		ADCH:   0x24, // ADC Data Register  Bytes
	}

	// Analog Comparator
	AC = struct {
		ACSR __reg
	}{
		ACSR: 0x28, // Analog Comparator Control And Status Register
	}

	// Two Wire Serial Interface
	TWI = struct {
		TWBR __reg
		TWCR __reg
		TWSR __reg
		TWDR __reg
		TWAR __reg
	}{
		TWBR: 0x20, // TWI Bit Rate register
		TWCR: 0x56, // TWI Control Register
		TWSR: 0x21, // TWI Status Register
		TWDR: 0x23, // TWI Data register
		TWAR: 0x22, // TWI (Slave) Address register
	}

	// USART
	USART = struct {
		UDR   __reg
		UCSRA __reg
		UCSRB __reg
		UCSRC __reg
		UBRRH __reg
		UBRRL __reg
	}{
		UDR:   0x2c, // USART I/O Data Register
		UCSRA: 0x2b, // USART Control and Status Register A
		UCSRB: 0x2a, // USART Control and Status Register B
		UCSRC: 0x40, // USART Control and Status Register C
		UBRRH: 0x40, // USART Baud Rate Register High Byte
		UBRRL: 0x29, // USART Baud Rate Register Low Byte
	}

	// I/O Port
	PORT = struct {
		PORTA __reg
		DDRA  __reg
		PINA  __reg
		PORTB __reg
		DDRB  __reg
		PINB  __reg
		PORTC __reg
		DDRC  __reg
		PINC  __reg
		PORTD __reg
		DDRD  __reg
		PIND  __reg
	}{
		PORTA: 0x3b, // Port A Data Register
		DDRA:  0x3a, // Port A Data Direction Register
		PINA:  0x39, // Port A Input Pins
		PORTB: 0x38, // Port B Data Register
		DDRB:  0x37, // Port B Data Direction Register
		PINB:  0x36, // Port B Input Pins
		PORTC: 0x35, // Port C Data Register
		DDRC:  0x34, // Port C Data Direction Register
		PINC:  0x33, // Port C Input Pins
		PORTD: 0x32, // Port D Data Register
		DDRD:  0x31, // Port D Data Direction Register
		PIND:  0x30, // Port D Input Pins
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPDR __reg
		SPSR __reg
		SPCR __reg
	}{
		SPDR: 0x2f, // SPI Data Register
		SPSR: 0x2e, // SPI Status Register
		SPCR: 0x2d, // SPI Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x3e, // EEPROM Address Register  Bytes
		EEARH: 0x3e, // EEPROM Address Register  Bytes
		EEDR:  0x3d, // EEPROM Data Register
		EECR:  0x3c, // EEPROM Control Register
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		TCCR0 __reg
		TCNT0 __reg
		OCR0  __reg
	}{
		TCCR0: 0x53, // Timer/Counter Control Register
		TCNT0: 0x52, // Timer/Counter Register
		OCR0:  0x5c, // Output Compare Register
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TCCR1A __reg
		TCCR1B __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		ICR1L  __reg
		ICR1H  __reg
	}{
		TCCR1A: 0x4f, // Timer/Counter1 Control Register A
		TCCR1B: 0x4e, // Timer/Counter1 Control Register B
		TCNT1L: 0x4c, // Timer/Counter1  Bytes
		TCNT1H: 0x4c, // Timer/Counter1  Bytes
		OCR1AL: 0x4a, // Timer/Counter1 Output Compare Register  Bytes
		OCR1AH: 0x4a, // Timer/Counter1 Output Compare Register  Bytes
		OCR1BL: 0x48, // Timer/Counter1 Output Compare Register  Bytes
		OCR1BH: 0x48, // Timer/Counter1 Output Compare Register  Bytes
		ICR1L:  0x46, // Timer/Counter1 Input Capture Register  Bytes
		ICR1H:  0x46, // Timer/Counter1 Input Capture Register  Bytes
	}

	// Timer/Counter, 8-bit Async
	TC8_ASYNC = struct {
		TCCR2 __reg
		TCNT2 __reg
		OCR2  __reg
		ASSR  __reg
	}{
		TCCR2: 0x45, // Timer/Counter2 Control Register
		TCNT2: 0x44, // Timer/Counter2
		OCR2:  0x43, // Timer/Counter2 Output Compare Register
		ASSR:  0x42, // Asynchronous Status Register
	}

	// External Interrupts
	EXINT = struct {
		GICR __reg
		GIFR __reg
	}{
		GICR: 0x5b, // General Interrupt Control Register
		GIFR: 0x5a, // General Interrupt Flag Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCR __reg
	}{
		WDTCR: 0x41, // Watchdog Timer Control Register
	}

	// CPU Registers
	CPU = struct {
		SREG   __reg
		SPL    __reg
		SPH    __reg
		OSCCAL __reg
		SPMCR  __reg
	}{
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		OSCCAL: 0x51, // Oscillator Calibration Value
		SPMCR:  0x57,
	}
)

// Bitfields for FUSE: Fuses
const (
	// HIGH
	HIGH_S8535C  = 0x80 // AT90S4434/8535 compatibility mode
	HIGH_WDTON   = 0x40 // Watch-dog Timer always on
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ  = 0x6  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled
	HIGH_CKOPT   = 0x10 // CKOPT fuse (operation dependent of CKSEL fuses)

	// LOW
	LOW_BODLEVEL  = 0x80 // Brownout detector trigger level
	LOW_BODEN     = 0x40 // Brown-out detection enabled
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX   = 0x1f // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS  = 0x7  // ADC Prescaler Select Bits
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD  = 0x80 // Analog Comparator Disable
	ACSR_ACBG = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO  = 0x20 // Analog Compare Output
	ACSR_ACI  = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS = 0x3  // Analog Comparator Interrupt Mode Select bits
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWCR: TWI Control Register
	TWCR_TWINT = 0x80 // TWI Interrupt Flag
	TWCR_TWEA  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN  = 0x4  // TWI Enable Bit
	TWCR_TWIE  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS  = 0xf8 // TWI Status
	TWSR_TWPS = 0x3  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA   = 0xfe // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for USART: USART
const (
	// UCSRA: USART Control and Status Register A
	UCSRA_RXC  = 0x80 // USART Receive Complete
	UCSRA_TXC  = 0x40 // USART Transmitt Complete
	UCSRA_UDRE = 0x20 // USART Data Register Empty
	UCSRA_FE   = 0x10 // Framing Error
	UCSRA_DOR  = 0x8  // Data overRun
	UCSRA_UPE  = 0x4  // Parity Error
	UCSRA_U2X  = 0x2  // Double the USART transmission speed
	UCSRA_MPCM = 0x1  // Multi-processor Communication Mode

	// UCSRB: USART Control and Status Register B
	UCSRB_RXCIE = 0x80 // RX Complete Interrupt Enable
	UCSRB_TXCIE = 0x40 // TX Complete Interrupt Enable
	UCSRB_UDRIE = 0x20 // USART Data register Empty Interrupt Enable
	UCSRB_RXEN  = 0x10 // Receiver Enable
	UCSRB_TXEN  = 0x8  // Transmitter Enable
	UCSRB_UCSZ2 = 0x4  // Character Size Bit 2
	UCSRB_RXB8  = 0x2  // Receive Data Bit 8
	UCSRB_TXB8  = 0x1  // Transmit Data Bit 8

	// UCSRC: USART Control and Status Register C
	UCSRC_URSEL = 0x80 // Register Select
	UCSRC_UMSEL = 0x40 // USART Mode Select
	UCSRC_UPM   = 0x30 // Parity Mode Bits
	UCSRC_USBS  = 0x8  // Stop Bit Select
	UCSRC_UCSZ  = 0x6  // Character Size Bits
	UCSRC_UCPOL = 0x1  // Clock Polarity

	// UBRRH: USART Baud Rate Register High Byte
	UBRRH_URSEL = 0x80 // Register Select
	UBRRH_UBRR1 = 0xc  // USART Baud Rate Register bit 11
	UBRRH_UBRR  = 0x3  // USART Baud Rate Register bits
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR  = 0x3  // SPI Clock Rate Selects
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0: Timer/Counter Control Register
	TCCR0_FOC0  = 0x80 // Force Output Compare
	TCCR0_WGM00 = 0x40 // Waveform Generation Mode 0
	TCCR0_COM0  = 0x30 // Compare Match Output Modes
	TCCR0_WGM01 = 0x8  // Waveform Generation Mode 1
	TCCR0_CS0   = 0x7  // Clock Selects
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_FOC1A = 0x8  // Force Output Compare 1A
	TCCR1A_FOC1B = 0x4  // Force Output Compare 1B
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR2: Timer/Counter2 Control Register
	TCCR2_FOC2  = 0x80 // Force Output Compare
	TCCR2_WGM20 = 0x40 // Waveform Genration Mode
	TCCR2_COM2  = 0x30 // Compare Output Mode bits
	TCCR2_WGM21 = 0x8  // Waveform Generation Mode
	TCCR2_CS2   = 0x7  // Clock Select bits

	// ASSR: Asynchronous Status Register
	ASSR_AS2    = 0x8 // Asynchronous Timer/counter2
	ASSR_TCN2UB = 0x4 // Timer/Counter2 Update Busy
	ASSR_OCR2UB = 0x2 // Output Compare Register2 Update Busy
	ASSR_TCR2UB = 0x1 // Timer/counter Control Register2 Update Busy
)

// Bitfields for EXINT: External Interrupts
const (
	// GICR: General Interrupt Control Register
	GICR_INT0  = 0x40 // External Interrupt Request 0 Enable
	GICR_INT1  = 0x80 // External Interrupt Request 1 Enable
	GICR_INT2  = 0x20 // External Interrupt Request 2 Enable
	GICR_IVSEL = 0x2  // Interrupt Vector Select
	GICR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// GIFR: General Interrupt Flag Register
	GIFR_INTF  = 0xc0 // External Interrupt Flags
	GIFR_INTF2 = 0x20 // External Interrupt Flag 2
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP  = 0x7  // Watch Dog Timer Prescaler bits
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// SPMCR
	SPMCR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCR_RWWSRE = 0x10 // Read-While-Write Section Read Enable
	SPMCR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCR_PGWRT  = 0x4  // Page Write
	SPMCR_PGERS  = 0x2  // Page Erase
	SPMCR_SPMEN  = 0x1  // Store Program Memory Enable
)
