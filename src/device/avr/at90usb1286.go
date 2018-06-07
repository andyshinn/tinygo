// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from AT90USB1286.atdf, see http://packs.download.atmel.com/

// +build avr,at90usb1286

// Device information for the AT90USB1286.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "AT90USB1286"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_INT3         = 4  // External Interrupt Request 3
	IRQ_INT4         = 5  // External Interrupt Request 4
	IRQ_INT5         = 6  // External Interrupt Request 5
	IRQ_INT6         = 7  // External Interrupt Request 6
	IRQ_INT7         = 8  // External Interrupt Request 7
	IRQ_PCINT0       = 9  // Pin Change Interrupt Request 0
	IRQ_USB_GEN      = 10 // USB General Interrupt Request
	IRQ_USB_COM      = 11 // USB Endpoint/Pipe Interrupt Communication Request
	IRQ_WDT          = 12 // Watchdog Time-out Interrupt
	IRQ_TIMER2_COMPA = 13 // Timer/Counter2 Compare Match A
	IRQ_TIMER2_COMPB = 14 // Timer/Counter2 Compare Match B
	IRQ_TIMER2_OVF   = 15 // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 16 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 17 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 18 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_COMPC = 19 // Timer/Counter1 Compare Match C
	IRQ_TIMER1_OVF   = 20 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 21 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 22 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_OVF   = 23 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 24 // SPI Serial Transfer Complete
	IRQ_USART1_RX    = 25 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 26 // USART1 Data register Empty
	IRQ_USART1_TX    = 27 // USART1, Tx Complete
	IRQ_ANALOG_COMP  = 28 // Analog Comparator
	IRQ_ADC          = 29 // ADC Conversion Complete
	IRQ_EE_READY     = 30 // EEPROM Ready
	IRQ_TIMER3_CAPT  = 31 // Timer/Counter3 Capture Event
	IRQ_TIMER3_COMPA = 32 // Timer/Counter3 Compare Match A
	IRQ_TIMER3_COMPB = 33 // Timer/Counter3 Compare Match B
	IRQ_TIMER3_COMPC = 34 // Timer/Counter3 Compare Match C
	IRQ_TIMER3_OVF   = 35 // Timer/Counter3 Overflow
	IRQ_TWI          = 36 // 2-wire Serial Interface
	IRQ_SPM_READY    = 37 // Store Program Memory Read
	IRQ_max          = 37 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		EXTENDED __reg
		HIGH     __reg
		LOW      __reg
	}{
		EXTENDED: 0x2,
		HIGH:     0x1,
		LOW:      0x0,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
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
		PORTE __reg
		DDRE  __reg
		PINE  __reg
		PORTF __reg
		DDRF  __reg
		PINF  __reg
	}{
		PORTA: 0x22, // Port A Data Register
		DDRA:  0x21, // Port A Data Direction Register
		PINA:  0x20, // Port A Input Pins
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
		PORTE: 0x2e, // Data Register, Port E
		DDRE:  0x2d, // Data Direction Register, Port E
		PINE:  0x2c, // Input Pins, Port E
		PORTF: 0x31, // Data Register, Port F
		DDRF:  0x30, // Data Direction Register, Port F
		PINF:  0x2f, // Input Pins, Port F
	}

	// CPU Registers
	CPU = struct {
		SREG   __reg
		SPL    __reg
		SPH    __reg
		XMCRA  __reg
		XMCRB  __reg
		OSCCAL __reg
		CLKPR  __reg
		SMCR   __reg
		EIND   __reg
		RAMPZ  __reg
		GPIOR2 __reg
		GPIOR1 __reg
		GPIOR0 __reg
		PRR1   __reg
		PRR0   __reg
	}{
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		XMCRA:  0x74, // External Memory Control Register A
		XMCRB:  0x75, // External Memory Control Register B
		OSCCAL: 0x66, // Oscillator Calibration Value
		CLKPR:  0x61,
		SMCR:   0x53, // Sleep Mode Control Register
		EIND:   0x5c, // Extended Indirect Register
		RAMPZ:  0x5b, // RAM Page Z Select Register
		GPIOR2: 0x4b, // General Purpose IO Register 2
		GPIOR1: 0x4a, // General Purpose IO Register 1
		GPIOR0: 0x3e, // General Purpose IO Register 0
		PRR1:   0x65, // Power Reduction Register1
		PRR0:   0x64, // Power Reduction Register0
	}

	// Two Wire Serial Interface
	TWI = struct {
		TWAMR __reg
		TWBR  __reg
		TWCR  __reg
		TWSR  __reg
		TWDR  __reg
		TWAR  __reg
	}{
		TWAMR: 0xbd, // TWI (Slave) Address Mask Register
		TWBR:  0xb8, // TWI Bit Rate register
		TWCR:  0xbc, // TWI Control Register
		TWSR:  0xb9, // TWI Status Register
		TWDR:  0xbb, // TWI Data register
		TWAR:  0xba, // TWI (Slave) Address register
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPCR __reg
		SPSR __reg
		SPDR __reg
	}{
		SPCR: 0x4c, // SPI Control Register
		SPSR: 0x4d, // SPI Status Register
		SPDR: 0x4e, // SPI Data Register
	}

	// USART
	USART = struct {
		UDR1   __reg
		UCSR1A __reg
		UCSR1B __reg
		UCSR1C __reg
		UBRR1L __reg
		UBRR1H __reg
	}{
		UDR1:   0xce, // USART I/O Data Register
		UCSR1A: 0xc8, // USART Control and Status Register A
		UCSR1B: 0xc9, // USART Control and Status Register B
		UCSR1C: 0xca, // USART Control and Status Register C
		UBRR1L: 0xcc, // USART Baud Rate Register  Bytes
		UBRR1H: 0xcc, // USART Baud Rate Register  Bytes
	}

	// USB Device Registers
	USB_DEVICE = struct {
		UEINT   __reg
		UEBCHX  __reg
		UEBCLX  __reg
		UEDATX  __reg
		UEIENX  __reg
		UESTA1X __reg
		UESTA0X __reg
		UECFG1X __reg
		UECFG0X __reg
		UECONX  __reg
		UERST   __reg
		UENUM   __reg
		UEINTX  __reg
		UDMFN   __reg
		UDFNUML __reg
		UDFNUMH __reg
		UDADDR  __reg
		UDIEN   __reg
		UDINT   __reg
		UDCON   __reg
	}{
		UEINT:   0xf4,
		UEBCHX:  0xf3,
		UEBCLX:  0xf2,
		UEDATX:  0xf1,
		UEIENX:  0xf0,
		UESTA1X: 0xef,
		UESTA0X: 0xee,
		UECFG1X: 0xed,
		UECFG0X: 0xec,
		UECONX:  0xeb,
		UERST:   0xea,
		UENUM:   0xe9,
		UEINTX:  0xe8,
		UDMFN:   0xe6,
		UDFNUML: 0xe4,
		UDFNUMH: 0xe4,
		UDADDR:  0xe3,
		UDIEN:   0xe2,
		UDINT:   0xe1,
		UDCON:   0xe0,
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCSR __reg
	}{
		SPMCSR: 0x57, // Store Program Memory Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x41, // EEPROM Address Register Low Bytes
		EEARH: 0x41, // EEPROM Address Register Low Bytes
		EEDR:  0x40, // EEPROM Data Register
		EECR:  0x3f, // EEPROM Control Register
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		OCR0B  __reg
		OCR0A  __reg
		TCNT0  __reg
		TCCR0B __reg
		TCCR0A __reg
		TIMSK0 __reg
		TIFR0  __reg
	}{
		OCR0B:  0x48, // Timer/Counter0 Output Compare Register
		OCR0A:  0x47, // Timer/Counter0 Output Compare Register
		TCNT0:  0x46, // Timer/Counter0
		TCCR0B: 0x45, // Timer/Counter Control Register B
		TCCR0A: 0x44, // Timer/Counter  Control Register A
		TIMSK0: 0x6e, // Timer/Counter0 Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter0 Interrupt Flag register
	}

	// Timer/Counter, 8-bit Async
	TC8_ASYNC = struct {
		TIMSK2 __reg
		TIFR2  __reg
		TCCR2A __reg
		TCCR2B __reg
		TCNT2  __reg
		OCR2B  __reg
		OCR2A  __reg
		ASSR   __reg
	}{
		TIMSK2: 0x70, // Timer/Counter Interrupt Mask register
		TIFR2:  0x37, // Timer/Counter Interrupt Flag Register
		TCCR2A: 0xb0, // Timer/Counter2 Control Register A
		TCCR2B: 0xb1, // Timer/Counter2 Control Register B
		TCNT2:  0xb2, // Timer/Counter2
		OCR2B:  0xb4, // Timer/Counter2 Output Compare Register B
		OCR2A:  0xb3, // Timer/Counter2 Output Compare Register A
		ASSR:   0xb6, // Asynchronous Status Register
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TCCR3A __reg
		TCCR3B __reg
		TCCR3C __reg
		TCNT3L __reg
		TCNT3H __reg
		OCR3AL __reg
		OCR3AH __reg
		OCR3BL __reg
		OCR3BH __reg
		OCR3CL __reg
		OCR3CH __reg
		ICR3L  __reg
		ICR3H  __reg
		TIMSK3 __reg
		TIFR3  __reg
		TCCR1A __reg
		TCCR1B __reg
		TCCR1C __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		OCR1CL __reg
		OCR1CH __reg
		ICR1L  __reg
		ICR1H  __reg
		TIMSK1 __reg
		TIFR1  __reg
	}{
		TCCR3A: 0x90, // Timer/Counter3 Control Register A
		TCCR3B: 0x91, // Timer/Counter3 Control Register B
		TCCR3C: 0x92, // Timer/Counter 3 Control Register C
		TCNT3L: 0x94, // Timer/Counter3  Bytes
		TCNT3H: 0x94, // Timer/Counter3  Bytes
		OCR3AL: 0x98, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3AH: 0x98, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3BL: 0x9a, // Timer/Counter3 Output Compare Register B  Bytes
		OCR3BH: 0x9a, // Timer/Counter3 Output Compare Register B  Bytes
		OCR3CL: 0x9c, // Timer/Counter3 Output Compare Register B  Bytes
		OCR3CH: 0x9c, // Timer/Counter3 Output Compare Register B  Bytes
		ICR3L:  0x96, // Timer/Counter3 Input Capture Register  Bytes
		ICR3H:  0x96, // Timer/Counter3 Input Capture Register  Bytes
		TIMSK3: 0x71, // Timer/Counter3 Interrupt Mask Register
		TIFR3:  0x38, // Timer/Counter3 Interrupt Flag register
		TCCR1A: 0x80, // Timer/Counter1 Control Register A
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1C: 0x82, // Timer/Counter 1 Control Register C
		TCNT1L: 0x84, // Timer/Counter1  Bytes
		TCNT1H: 0x84, // Timer/Counter1  Bytes
		OCR1AL: 0x88, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1AH: 0x88, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1BL: 0x8a, // Timer/Counter1 Output Compare Register B  Bytes
		OCR1BH: 0x8a, // Timer/Counter1 Output Compare Register B  Bytes
		OCR1CL: 0x8c, // Timer/Counter1 Output Compare Register C  Bytes
		OCR1CH: 0x8c, // Timer/Counter1 Output Compare Register C  Bytes
		ICR1L:  0x86, // Timer/Counter1 Input Capture Register  Bytes
		ICR1H:  0x86, // Timer/Counter1 Input Capture Register  Bytes
		TIMSK1: 0x6f, // Timer/Counter1 Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter1 Interrupt Flag register
	}

	// JTAG Interface
	JTAG = struct {
		OCDR __reg
	}{
		OCDR: 0x51, // On-Chip Debug Related Register in I/O Memory
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EICRB  __reg
		EIMSK  __reg
		EIFR   __reg
		PCMSK0 __reg
		PCIFR  __reg
		PCICR  __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register A
		EICRB:  0x6a, // External Interrupt Control Register B
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCMSK0: 0x6b, // Pin Change Mask Register 0
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
	}

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX  __reg
		ADCSRA __reg
		ADCL   __reg
		ADCH   __reg
		DIDR0  __reg
	}{
		ADMUX:  0x7c, // The ADC multiplexer Selection Register
		ADCSRA: 0x7a, // The ADC Control and Status register
		ADCL:   0x78, // ADC Data Register  Bytes
		ADCH:   0x78, // ADC Data Register  Bytes
		DIDR0:  0x7e, // Digital Input Disable Register 1
	}

	// Analog Comparator
	AC = struct {
		ACSR  __reg
		DIDR1 __reg
	}{
		ACSR:  0x50, // Analog Comparator Control And Status Register
		DIDR1: 0x7f,
	}

	// Phase Locked Loop
	PLL = struct {
		PLLCSR __reg
	}{
		PLLCSR: 0x49, // PLL Status and Control register
	}

	// USB Controller
	USB_GLOBAL = struct {
		USBINT __reg
		USBSTA __reg
		USBCON __reg
		UHWCON __reg
	}{
		USBINT: 0xda,
		USBSTA: 0xd9,
		USBCON: 0xd8, // USB General Control Register
		UHWCON: 0xd7, // USB Hardware Configuration Register
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL = 0x7 // Brown-out Detector trigger level
	EXTENDED_HWBE     = 0x8 // Hardware Boot Enable

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON   = 0x10 // Watchdog timer always on
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ  = 0x6  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTC7
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
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

	// XMCRA: External Memory Control Register A
	XMCRA_SRE  = 0x80 // External SRAM Enable
	XMCRA_SRL  = 0x70 // Wait state page limit
	XMCRA_SRW1 = 0xc  // Wait state select bit upper page
	XMCRA_SRW0 = 0x3  // Wait state select bit lower page

	// XMCRB: External Memory Control Register B
	XMCRB_XMBK = 0x80 // External Memory Bus Keeper Enable
	XMCRB_XMM  = 0x7  // External Memory High Mask

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// CLKPR
	CLKPR_CLKPCE = 0x80
	CLKPR_CLKPS  = 0xf

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select bits
	SMCR_SE = 0x1 // Sleep Enable

	// GPIOR2: General Purpose IO Register 2
	GPIOR2_GPIOR = 0xff // General Purpose IO Register 2 bis

	// GPIOR1: General Purpose IO Register 1
	GPIOR1_GPIOR = 0xff // General Purpose IO Register 1 bis

	// GPIOR0: General Purpose IO Register 0
	GPIOR0_GPIOR07 = 0x80 // General Purpose IO Register 0 bit 7
	GPIOR0_GPIOR06 = 0x40 // General Purpose IO Register 0 bit 6
	GPIOR0_GPIOR05 = 0x20 // General Purpose IO Register 0 bit 5
	GPIOR0_GPIOR04 = 0x10 // General Purpose IO Register 0 bit 4
	GPIOR0_GPIOR03 = 0x8  // General Purpose IO Register 0 bit 3
	GPIOR0_GPIOR02 = 0x4  // General Purpose IO Register 0 bit 2
	GPIOR0_GPIOR01 = 0x2  // General Purpose IO Register 0 bit 1
	GPIOR0_GPIOR00 = 0x1  // General Purpose IO Register 0 bit 0

	// PRR1: Power Reduction Register1
	PRR1_PRUSB    = 0x80 // Power Reduction USB
	PRR1_PRTIM3   = 0x8  // Power Reduction Timer/Counter3
	PRR1_PRUSART1 = 0x1  // Power Reduction USART1

	// PRR0: Power Reduction Register0
	PRR0_PRTWI  = 0x80 // Power Reduction TWI
	PRR0_PRTIM2 = 0x40 // Power Reduction Timer/Counter2
	PRR0_PRTIM0 = 0x20 // Power Reduction Timer/Counter0
	PRR0_PRTIM1 = 0x8  // Power Reduction Timer/Counter1
	PRR0_PRSPI  = 0x4  // Power Reduction Serial Peripheral Interface
	PRR0_PRADC  = 0x1  // Power Reduction ADC
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM = 0xfe

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

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR  = 0x3  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit
)

// Bitfields for USART: USART
const (
	// UCSR1A: USART Control and Status Register A
	UCSR1A_RXC1  = 0x80 // USART Receive Complete
	UCSR1A_TXC1  = 0x40 // USART Transmitt Complete
	UCSR1A_UDRE1 = 0x20 // USART Data Register Empty
	UCSR1A_FE1   = 0x10 // Framing Error
	UCSR1A_DOR1  = 0x8  // Data overRun
	UCSR1A_UPE1  = 0x4  // Parity Error
	UCSR1A_U2X1  = 0x2  // Double the USART transmission speed
	UCSR1A_MPCM1 = 0x1  // Multi-processor Communication Mode

	// UCSR1B: USART Control and Status Register B
	UCSR1B_RXCIE1 = 0x80 // RX Complete Interrupt Enable
	UCSR1B_TXCIE1 = 0x40 // TX Complete Interrupt Enable
	UCSR1B_UDRIE1 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR1B_RXEN1  = 0x10 // Receiver Enable
	UCSR1B_TXEN1  = 0x8  // Transmitter Enable
	UCSR1B_UCSZ12 = 0x4  // Character Size
	UCSR1B_RXB81  = 0x2  // Receive Data Bit 8
	UCSR1B_TXB81  = 0x1  // Transmit Data Bit 8

	// UCSR1C: USART Control and Status Register C
	UCSR1C_UMSEL1 = 0xc0 // USART Mode Select
	UCSR1C_UPM1   = 0x30 // Parity Mode Bits
	UCSR1C_USBS1  = 0x8  // Stop Bit Select
	UCSR1C_UCSZ1  = 0x6  // Character Size
	UCSR1C_UCPOL1 = 0x1  // Clock Polarity
)

// Bitfields for USB_DEVICE: USB Device Registers
const (
	// UEIENX
	UEIENX_FLERRE   = 0x80
	UEIENX_NAKINE   = 0x40
	UEIENX_NAKOUTE  = 0x10
	UEIENX_RXSTPE   = 0x8
	UEIENX_RXOUTE   = 0x4
	UEIENX_STALLEDE = 0x2
	UEIENX_TXINE    = 0x1

	// UESTA1X
	UESTA1X_CTRLDIR = 0x4
	UESTA1X_CURRBK  = 0x3

	// UESTA0X
	UESTA0X_CFGOK   = 0x80
	UESTA0X_OVERFI  = 0x40
	UESTA0X_UNDERFI = 0x20
	UESTA0X_DTSEQ   = 0xc
	UESTA0X_NBUSYBK = 0x3

	// UECFG1X
	UECFG1X_EPSIZE = 0x70
	UECFG1X_EPBK   = 0xc
	UECFG1X_ALLOC  = 0x2

	// UECFG0X
	UECFG0X_EPTYPE = 0xc0
	UECFG0X_EPDIR  = 0x1

	// UECONX
	UECONX_STALLRQ  = 0x20
	UECONX_STALLRQC = 0x10
	UECONX_RSTDT    = 0x8
	UECONX_EPEN     = 0x1

	// UERST
	UERST_EPRST = 0x7f

	// UEINTX
	UEINTX_FIFOCON  = 0x80
	UEINTX_NAKINI   = 0x40
	UEINTX_RWAL     = 0x20
	UEINTX_NAKOUTI  = 0x10
	UEINTX_RXSTPI   = 0x8
	UEINTX_RXOUTI   = 0x4
	UEINTX_STALLEDI = 0x2
	UEINTX_TXINI    = 0x1

	// UDMFN
	UDMFN_FNCERR = 0x10

	// UDADDR
	UDADDR_ADDEN = 0x80
	UDADDR_UADD  = 0x7f

	// UDIEN
	UDIEN_UPRSME  = 0x40
	UDIEN_EORSME  = 0x20
	UDIEN_WAKEUPE = 0x10
	UDIEN_EORSTE  = 0x8
	UDIEN_SOFE    = 0x4
	UDIEN_SUSPE   = 0x1

	// UDINT
	UDINT_UPRSMI  = 0x40
	UDINT_EORSMI  = 0x20
	UDINT_WAKEUPI = 0x10
	UDINT_EORSTI  = 0x8
	UDINT_SOFI    = 0x4
	UDINT_SUSPI   = 0x1

	// UDCON
	UDCON_LSM    = 0x4
	UDCON_RMWKUP = 0x2
	UDCON_DETACH = 0x1
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM  = 0x30 // EEPROM Programming Mode Bits
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE = 0x4  // EEPROM Master Write Enable
	EECR_EEPE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS0   = 0x7 // Clock Select

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A = 0xc0 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B = 0x30 // Compare Output Mode, Fast PWm
	TCCR0A_WGM0  = 0x3  // Waveform Generation Mode

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TIMSK2: Timer/Counter Interrupt Mask register
	TIMSK2_OCIE2B = 0x4 // Timer/Counter2 Output Compare Match B Interrupt Enable
	TIMSK2_OCIE2A = 0x2 // Timer/Counter2 Output Compare Match A Interrupt Enable
	TIMSK2_TOIE2  = 0x1 // Timer/Counter2 Overflow Interrupt Enable

	// TIFR2: Timer/Counter Interrupt Flag Register
	TIFR2_OCF2B = 0x4 // Output Compare Flag 2B
	TIFR2_OCF2A = 0x2 // Output Compare Flag 2A
	TIFR2_TOV2  = 0x1 // Timer/Counter2 Overflow Flag

	// TCCR2A: Timer/Counter2 Control Register A
	TCCR2A_COM2A = 0xc0 // Compare Output Mode bits
	TCCR2A_COM2B = 0x30 // Compare Output Mode bits
	TCCR2A_WGM2  = 0x3  // Waveform Genration Mode

	// TCCR2B: Timer/Counter2 Control Register B
	TCCR2B_FOC2A = 0x80 // Force Output Compare A
	TCCR2B_FOC2B = 0x40 // Force Output Compare B
	TCCR2B_WGM22 = 0x8  // Waveform Generation Mode
	TCCR2B_CS2   = 0x7  // Clock Select bits

	// ASSR: Asynchronous Status Register
	ASSR_EXCLK   = 0x40 // Enable External Clock Input
	ASSR_AS2     = 0x20 // Asynchronous Timer/Counter2
	ASSR_TCN2UB  = 0x10 // Timer/Counter2 Update Busy
	ASSR_OCR2AUB = 0x8  // Output Compare Register2 Update Busy
	ASSR_OCR2BUB = 0x4  // Output Compare Register 2 Update Busy
	ASSR_TCR2AUB = 0x2  // Timer/Counter Control Register2 Update Busy
	ASSR_TCR2BUB = 0x1  // Timer/Counter Control Register2 Update Busy
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR3A: Timer/Counter3 Control Register A
	TCCR3A_COM3A = 0xc0 // Compare Output Mode 1A, bits
	TCCR3A_COM3B = 0x30 // Compare Output Mode 3B, bits
	TCCR3A_COM3C = 0xc  // Compare Output Mode 3C, bits
	TCCR3A_WGM3  = 0x3  // Waveform Generation Mode

	// TCCR3B: Timer/Counter3 Control Register B
	TCCR3B_ICNC3 = 0x80 // Input Capture 3 Noise Canceler
	TCCR3B_ICES3 = 0x40 // Input Capture 3 Edge Select
	TCCR3B_WGM3  = 0x18 // Waveform Generation Mode
	TCCR3B_CS3   = 0x7  // Prescaler source of Timer/Counter 3

	// TCCR3C: Timer/Counter 3 Control Register C
	TCCR3C_FOC3A = 0x80 // Force Output Compare 3A
	TCCR3C_FOC3B = 0x40 // Force Output Compare 3B
	TCCR3C_FOC3C = 0x20 // Force Output Compare 3C

	// TIMSK3: Timer/Counter3 Interrupt Mask Register
	TIMSK3_ICIE3  = 0x20 // Timer/Counter3 Input Capture Interrupt Enable
	TIMSK3_OCIE3C = 0x8  // Timer/Counter3 Output Compare C Match Interrupt Enable
	TIMSK3_OCIE3B = 0x4  // Timer/Counter3 Output Compare B Match Interrupt Enable
	TIMSK3_OCIE3A = 0x2  // Timer/Counter3 Output Compare A Match Interrupt Enable
	TIMSK3_TOIE3  = 0x1  // Timer/Counter3 Overflow Interrupt Enable

	// TIFR3: Timer/Counter3 Interrupt Flag register
	TIFR3_ICF3  = 0x20 // Input Capture Flag 3
	TIFR3_OCF3C = 0x8  // Output Compare Flag 3C
	TIFR3_OCF3B = 0x4  // Output Compare Flag 3B
	TIFR3_OCF3A = 0x2  // Output Compare Flag 3A
	TIFR3_TOV3  = 0x1  // Timer/Counter3 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_COM1C = 0xc  // Compare Output Mode 1C, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter 1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare 1A
	TCCR1C_FOC1B = 0x40 // Force Output Compare 1B
	TCCR1C_FOC1C = 0x20 // Force Output Compare 1C

	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1C = 0x8  // Timer/Counter1 Output Compare C Match Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1C = 0x8  // Output Compare Flag 1C
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC3 = 0xc0 // External Interrupt Sense Control Bit
	EICRA_ISC2 = 0x30 // External Interrupt Sense Control Bit
	EICRA_ISC1 = 0xc  // External Interrupt Sense Control Bit
	EICRA_ISC0 = 0x3  // External Interrupt Sense Control Bit

	// EICRB: External Interrupt Control Register B
	EICRB_ISC7 = 0xc0 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC6 = 0x30 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC5 = 0xc  // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC4 = 0x3  // External Interrupt 7-4 Sense Control Bit

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT = 0xff // External Interrupt Request 7 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF = 0xff // External Interrupt Flags

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flag 0

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enable 0
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
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS  = 0x7  // ADC Prescaler Select Bits

	// DIDR0: Digital Input Disable Register 1
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable
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

	// DIDR1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
)

// Bitfields for PLL: Phase Locked Loop
const (
	// PLLCSR: PLL Status and Control register
	PLLCSR_PLLP  = 0x1c // PLL prescaler Bits
	PLLCSR_PLLE  = 0x2  // PLL Enable Bit
	PLLCSR_PLOCK = 0x1  // PLL Lock Status Bit
)

// Bitfields for USB_GLOBAL: USB Controller
const (
	// USBINT
	USBINT_IDTI   = 0x2
	USBINT_VBUSTI = 0x1

	// USBSTA
	USBSTA_SPEED = 0x8
	USBSTA_ID    = 0x2
	USBSTA_VBUS  = 0x1

	// USBCON: USB General Control Register
	USBCON_USBE    = 0x80
	USBCON_HOST    = 0x40
	USBCON_FRZCLK  = 0x20
	USBCON_OTGPADE = 0x10
	USBCON_IDTE    = 0x2
	USBCON_VBUSTE  = 0x1

	// UHWCON: USB Hardware Configuration Register
	UHWCON_UIMOD  = 0x80
	UHWCON_UIDE   = 0x40
	UHWCON_UVCONE = 0x10
	UHWCON_UVREGE = 0x1
)
