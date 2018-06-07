// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega128.atdf, see http://packs.download.atmel.com/

// +build avr,atmega128

// Device information for the ATmega128.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega128"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset, Watchdog Reset and JTAG AVR Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_INT3         = 4  // External Interrupt Request 3
	IRQ_INT4         = 5  // External Interrupt Request 4
	IRQ_INT5         = 6  // External Interrupt Request 5
	IRQ_INT6         = 7  // External Interrupt Request 6
	IRQ_INT7         = 8  // External Interrupt Request 7
	IRQ_TIMER2_COMP  = 9  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 10 // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 11 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 12 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 13 // Timer/Counter Compare Match B
	IRQ_TIMER1_OVF   = 14 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMP  = 15 // Timer/Counter0 Compare Match
	IRQ_TIMER0_OVF   = 16 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 17 // SPI Serial Transfer Complete
	IRQ_USART0_RX    = 18 // USART0, Rx Complete
	IRQ_USART0_UDRE  = 19 // USART0 Data Register Empty
	IRQ_USART0_TX    = 20 // USART0, Tx Complete
	IRQ_ADC          = 21 // ADC Conversion Complete
	IRQ_EE_READY     = 22 // EEPROM Ready
	IRQ_ANALOG_COMP  = 23 // Analog Comparator
	IRQ_TIMER1_COMPC = 24 // Timer/Counter1 Compare Match C
	IRQ_TIMER3_CAPT  = 25 // Timer/Counter3 Capture Event
	IRQ_TIMER3_COMPA = 26 // Timer/Counter3 Compare Match A
	IRQ_TIMER3_COMPB = 27 // Timer/Counter3 Compare Match B
	IRQ_TIMER3_COMPC = 28 // Timer/Counter3 Compare Match C
	IRQ_TIMER3_OVF   = 29 // Timer/Counter3 Overflow
	IRQ_USART1_RX    = 30 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 31 // USART1, Data Register Empty
	IRQ_USART1_TX    = 32 // USART1, Tx Complete
	IRQ_TWI          = 33 // 2-wire Serial Interface
	IRQ_SPM_READY    = 34 // Store Program Memory Read
	IRQ_max          = 34 // Highest interrupt number on this device.
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

	// Analog Comparator
	AC = struct {
		ACSR __reg
	}{
		ACSR: 0x28, // Analog Comparator Control And Status Register
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

	// Two Wire Serial Interface
	TWI = struct {
		TWBR __reg
		TWCR __reg
		TWSR __reg
		TWDR __reg
		TWAR __reg
	}{
		TWBR: 0x70, // TWI Bit Rate register
		TWCR: 0x74, // TWI Control Register
		TWSR: 0x71, // TWI Status Register
		TWDR: 0x73, // TWI Data register
		TWAR: 0x72, // TWI (Slave) Address register
	}

	// USART
	USART = struct {
		UDR0   __reg
		UCSR0A __reg
		UCSR0B __reg
		UCSR0C __reg
		UBRR0H __reg
		UBRR0L __reg
		UDR1   __reg
		UCSR1A __reg
		UCSR1B __reg
		UCSR1C __reg
		UBRR1H __reg
		UBRR1L __reg
	}{
		UDR0:   0x2c, // USART I/O Data Register
		UCSR0A: 0x2b, // USART Control and Status Register A
		UCSR0B: 0x2a, // USART Control and Status Register B
		UCSR0C: 0x95, // USART Control and Status Register C
		UBRR0H: 0x90, // USART Baud Rate Register Hight Byte
		UBRR0L: 0x29, // USART Baud Rate Register Low Byte
		UDR1:   0x9c, // USART I/O Data Register
		UCSR1A: 0x9b, // USART Control and Status Register A
		UCSR1B: 0x9a, // USART Control and Status Register B
		UCSR1C: 0x9d, // USART Control and Status Register C
		UBRR1H: 0x98, // USART Baud Rate Register Hight Byte
		UBRR1L: 0x99, // USART Baud Rate Register Low Byte
	}

	// CPU Registers
	CPU = struct {
		SREG   __reg
		SPL    __reg
		SPH    __reg
		MCUCR  __reg
		XDIV   __reg
		XMCRA  __reg
		XMCRB  __reg
		OSCCAL __reg
		RAMPZ  __reg
	}{
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		MCUCR:  0x55, // MCU Control Register
		XDIV:   0x5c, // XTAL Divide Control Register
		XMCRA:  0x6d, // External Memory Control Register A
		XMCRB:  0x6c, // External Memory Control Register B
		OSCCAL: 0x6f, // Oscillator Calibration Value
		RAMPZ:  0x5b, // RAM Page Z Select Register
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCSR __reg
	}{
		SPMCSR: 0x68, // Store Program Memory Control Register
	}

	// JTAG Interface
	JTAG = struct {
		OCDR __reg
	}{
		OCDR: 0x42, // On-Chip Debug Related Register in I/O Memory
	}

	// Other Registers
	MISC = struct {
	}{}

	// External Interrupts
	EXINT = struct {
		EICRA __reg
		EICRB __reg
		EIMSK __reg
		EIFR  __reg
	}{
		EICRA: 0x6a, // External Interrupt Control Register A
		EICRB: 0x5a, // External Interrupt Control Register B
		EIMSK: 0x59, // External Interrupt Mask Register
		EIFR:  0x58, // External Interrupt Flag Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x3e, // EEPROM Read/Write Access  Bytes
		EEARH: 0x3e, // EEPROM Read/Write Access  Bytes
		EEDR:  0x3d, // EEPROM Data Register
		EECR:  0x3c, // EEPROM Control Register
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
		PORTG __reg
		DDRG  __reg
		PING  __reg
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
		PORTE: 0x23, // Data Register, Port E
		DDRE:  0x22, // Data Direction Register, Port E
		PINE:  0x21, // Input Pins, Port E
		PORTF: 0x62, // Data Register, Port F
		DDRF:  0x61, // Data Direction Register, Port F
		PINF:  0x20, // Input Pins, Port F
		PORTG: 0x65, // Data Register, Port G
		DDRG:  0x64, // Data Direction Register, Port G
		PING:  0x63, // Input Pins, Port G
	}

	// Timer/Counter, 8-bit Async
	TC8_ASYNC = struct {
		TCCR0 __reg
		TCNT0 __reg
		OCR0  __reg
		ASSR  __reg
	}{
		TCCR0: 0x53, // Timer/Counter Control Register
		TCNT0: 0x52, // Timer/Counter Register
		OCR0:  0x51, // Output Compare Register
		ASSR:  0x50, // Asynchronus Status Register
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
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
	}{
		TCCR1A: 0x4f, // Timer/Counter1 Control Register A
		TCCR1B: 0x4e, // Timer/Counter1 Control Register B
		TCCR1C: 0x7a, // Timer/Counter1 Control Register C
		TCNT1L: 0x4c, // Timer/Counter1  Bytes
		TCNT1H: 0x4c, // Timer/Counter1  Bytes
		OCR1AL: 0x4a, // Timer/Counter1 Output Compare Register  Bytes
		OCR1AH: 0x4a, // Timer/Counter1 Output Compare Register  Bytes
		OCR1BL: 0x48, // Timer/Counter1 Output Compare Register  Bytes
		OCR1BH: 0x48, // Timer/Counter1 Output Compare Register  Bytes
		OCR1CL: 0x78, // Timer/Counter1 Output Compare Register  Bytes
		OCR1CH: 0x78, // Timer/Counter1 Output Compare Register  Bytes
		ICR1L:  0x46, // Timer/Counter1 Input Capture Register  Bytes
		ICR1H:  0x46, // Timer/Counter1 Input Capture Register  Bytes
		TCCR3A: 0x8b, // Timer/Counter3 Control Register A
		TCCR3B: 0x8a, // Timer/Counter3 Control Register B
		TCCR3C: 0x8c, // Timer/Counter3 Control Register C
		TCNT3L: 0x88, // Timer/Counter3  Bytes
		TCNT3H: 0x88, // Timer/Counter3  Bytes
		OCR3AL: 0x86, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3AH: 0x86, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3BL: 0x84, // Timer/Counter3 Output Compare Register B  Bytes
		OCR3BH: 0x84, // Timer/Counter3 Output Compare Register B  Bytes
		OCR3CL: 0x82, // Timer/Counter3 Output compare Register C  Bytes
		OCR3CH: 0x82, // Timer/Counter3 Output compare Register C  Bytes
		ICR3L:  0x80, // Timer/Counter3 Input Capture Register  Bytes
		ICR3H:  0x80, // Timer/Counter3 Input Capture Register  Bytes
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		TCCR2 __reg
		TCNT2 __reg
		OCR2  __reg
	}{
		TCCR2: 0x45, // Timer/Counter Control Register
		TCNT2: 0x44, // Timer/Counter Register
		OCR2:  0x43, // Output Compare Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCR __reg
	}{
		WDTCR: 0x41, // Watchdog Timer Control Register
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
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_M103C = 0x2 // ATmega103 Compatibility Mode
	EXTENDED_WDTON = 0x1 // Watchdog Timer always on

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
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
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmitt Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data overRun
	UCSR0A_UPE0  = 0x4  // Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART transmission speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_UMSEL0 = 0x40 // USART Mode Select
	UCSR0C_UPM0   = 0x30 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ0  = 0x6  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity

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
	UCSR1C_UMSEL1 = 0x40 // USART Mode Select
	UCSR1C_UPM1   = 0x30 // Parity Mode Bits
	UCSR1C_USBS1  = 0x8  // Stop Bit Select
	UCSR1C_UCSZ1  = 0x6  // Character Size
	UCSR1C_UCPOL1 = 0x1  // Clock Polarity
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

	// MCUCR: MCU Control Register
	MCUCR_SRE   = 0x80 // External SRAM Enable
	MCUCR_SRW10 = 0x40 // External SRAM Wait State Select
	MCUCR_SE    = 0x20 // Sleep Enable
	MCUCR_SM    = 0x18 // Sleep Mode Select
	MCUCR_SM2   = 0x4  // Sleep Mode Select
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// XDIV: XTAL Divide Control Register
	XDIV_XDIVEN = 0x80 // XTAL Divide Enable
	XDIV_XDIV6  = 0x40 // XTAL Divide Select Bit 6
	XDIV_XDIV5  = 0x20 // XTAL Divide Select Bit 5
	XDIV_XDIV4  = 0x10 // XTAL Divide Select Bit 4
	XDIV_XDIV3  = 0x8  // XTAL Divide Select Bit 3
	XDIV_XDIV2  = 0x4  // XTAL Divide Select Bit 2
	XDIV_XDIV1  = 0x2  // XTAL Divide Select Bit 1
	XDIV_XDIV0  = 0x1  // XTAL Divide Select Bit 0

	// XMCRA: External Memory Control Register A
	XMCRA_SRL   = 0x70 // Wait state page limit
	XMCRA_SRW0  = 0xc  // Wait state select bit lower page
	XMCRA_SRW11 = 0x2  // Wait state select bit upper page

	// XMCRB: External Memory Control Register B
	XMCRB_XMBK = 0x80 // External Memory Bus Keeper Enable
	XMCRB_XMM  = 0x7  // External Memory High Mask

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// RAMPZ: RAM Page Z Select Register
	RAMPZ_RAMPZ0 = 0x1 // RAM Page Z Select Register Bit 0
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for JTAG: JTAG Interface
const (
	// OCDR: On-Chip Debug Related Register in I/O Memory
	OCDR_OCDR = 0xff // On-Chip Debug Register Bits
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
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR0: Timer/Counter Control Register
	TCCR0_FOC0  = 0x80 // Force Output Compare
	TCCR0_WGM00 = 0x40 // Waveform Generation Mode 0
	TCCR0_COM0  = 0x30 // Compare Match Output Modes
	TCCR0_WGM01 = 0x8  // Waveform Generation Mode 1
	TCCR0_CS0   = 0x7  // Clock Selects

	// ASSR: Asynchronus Status Register
	ASSR_AS0    = 0x8 // Asynchronus Timer/Counter 0
	ASSR_TCN0UB = 0x4 // Timer/Counter0 Update Busy
	ASSR_OCR0UB = 0x2 // Output Compare register 0 Busy
	ASSR_TCR0UB = 0x1 // Timer/Counter Control Register 0 Update Busy
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_COM1C = 0xc  // Compare Output Mode 1C, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Clock Select1 bits

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare for channel A
	TCCR1C_FOC1B = 0x40 // Force Output Compare for channel B
	TCCR1C_FOC1C = 0x20 // Force Output Compare for channel C

	// TCCR3A: Timer/Counter3 Control Register A
	TCCR3A_COM3A = 0xc0 // Compare Output Mode 3A, bits
	TCCR3A_COM3B = 0x30 // Compare Output Mode 3B, bits
	TCCR3A_COM3C = 0xc  // Compare Output Mode 3C, bits
	TCCR3A_WGM3  = 0x3  // Waveform Generation Mode Bits

	// TCCR3B: Timer/Counter3 Control Register B
	TCCR3B_ICNC3 = 0x80 // Input Capture 3  Noise Canceler
	TCCR3B_ICES3 = 0x40 // Input Capture 3 Edge Select
	TCCR3B_WGM3  = 0x18 // Waveform Generation Mode
	TCCR3B_CS3   = 0x7  // Clock Select3 bits

	// TCCR3C: Timer/Counter3 Control Register C
	TCCR3C_FOC3A = 0x80 // Force Output Compare for channel A
	TCCR3C_FOC3B = 0x40 // Force Output Compare for channel B
	TCCR3C_FOC3C = 0x20 // Force Output Compare for channel C
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR2: Timer/Counter Control Register
	TCCR2_FOC2  = 0x80 // Force Output Compare
	TCCR2_WGM20 = 0x40 // Wafeform Generation Mode
	TCCR2_COM2  = 0x30 // Compare Match Output Mode
	TCCR2_WGM21 = 0x8  // Waveform Generation Mode
	TCCR2_CS2   = 0x7  // Clock Select
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP  = 0x7  // Watch Dog Timer Prescaler bits
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX   = 0x1f // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN = 0x80 // ADC Enable
	ADCSRA_ADSC = 0x40 // ADC Start Conversion
	ADCSRA_ADFR = 0x20 // ADC  Free Running Select
	ADCSRA_ADIF = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS = 0x7  // ADC  Prescaler Select Bits
)
