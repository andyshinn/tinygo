// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega328.atdf, see http://packs.download.atmel.com/

// +build avr,atmega328

// Device information for the ATmega328.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega328"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_PCINT0       = 3  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 4  // Pin Change Interrupt Request 1
	IRQ_PCINT2       = 5  // Pin Change Interrupt Request 2
	IRQ_WDT          = 6  // Watchdog Time-out Interrupt
	IRQ_TIMER2_COMPA = 7  // Timer/Counter2 Compare Match A
	IRQ_TIMER2_COMPB = 8  // Timer/Counter2 Compare Match B
	IRQ_TIMER2_OVF   = 9  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 10 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 11 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 12 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 13 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 14 // TimerCounter0 Compare Match A
	IRQ_TIMER0_COMPB = 15 // TimerCounter0 Compare Match B
	IRQ_TIMER0_OVF   = 16 // Timer/Couner0 Overflow
	IRQ_SPI_STC      = 17 // SPI Serial Transfer Complete
	IRQ_USART_RX     = 18 // USART Rx Complete
	IRQ_USART_UDRE   = 19 // USART, Data Register Empty
	IRQ_USART_TX     = 20 // USART Tx Complete
	IRQ_ADC          = 21 // ADC Conversion Complete
	IRQ_EE_READY     = 22 // EEPROM Ready
	IRQ_ANALOG_COMP  = 23 // Analog Comparator
	IRQ_TWI          = 24 // Two-wire Serial Interface
	IRQ_SPM_READY    = 25 // Store Program Memory Read
	IRQ_max          = 25 // Highest interrupt number on this device.
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

	// USART
	USART = struct {
		UDR0   __reg
		UCSR0A __reg
		UCSR0B __reg
		UCSR0C __reg
		UBRR0L __reg
		UBRR0H __reg
	}{
		UDR0:   0xc6, // USART I/O Data Register
		UCSR0A: 0xc0, // USART Control and Status Register A
		UCSR0B: 0xc1, // USART Control and Status Register B
		UCSR0C: 0xc2, // USART Control and Status Register C
		UBRR0L: 0xc4, // USART Baud Rate Register Bytes
		UBRR0H: 0xc4, // USART Baud Rate Register Bytes
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

	// Timer/Counter, 16-bit
	TC16 = struct {
		TIMSK1 __reg
		TIFR1  __reg
		TCCR1A __reg
		TCCR1B __reg
		TCCR1C __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		ICR1L  __reg
		ICR1H  __reg
	}{
		TIMSK1: 0x6f, // Timer/Counter Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter Interrupt Flag register
		TCCR1A: 0x80, // Timer/Counter1 Control Register A
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1C: 0x82, // Timer/Counter1 Control Register C
		TCNT1L: 0x84, // Timer/Counter1 Bytes
		TCNT1H: 0x84, // Timer/Counter1 Bytes
		OCR1AL: 0x88, // Timer/Counter1 Output Compare Register Bytes
		OCR1AH: 0x88, // Timer/Counter1 Output Compare Register Bytes
		OCR1BL: 0x8a, // Timer/Counter1 Output Compare Register Bytes
		OCR1BH: 0x8a, // Timer/Counter1 Output Compare Register Bytes
		ICR1L:  0x86, // Timer/Counter1 Input Capture Register Bytes
		ICR1H:  0x86, // Timer/Counter1 Input Capture Register Bytes
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

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX  __reg
		ADCL   __reg
		ADCH   __reg
		ADCSRA __reg
		ADCSRB __reg
		DIDR0  __reg
	}{
		ADMUX:  0x7c, // The ADC multiplexer Selection Register
		ADCL:   0x78, // ADC Data Register Bytes
		ADCH:   0x78, // ADC Data Register Bytes
		ADCSRA: 0x7a, // The ADC Control and Status register A
		ADCSRB: 0x7b, // The ADC Control and Status register B
		DIDR0:  0x7e, // Digital Input Disable Register
	}

	// Analog Comparator
	AC = struct {
		ACSR  __reg
		DIDR1 __reg
	}{
		ACSR:  0x50, // Analog Comparator Control And Status Register
		DIDR1: 0x7f, // Digital Input Disable Register 1
	}

	// I/O Port
	PORT = struct {
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
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
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

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EIMSK  __reg
		EIFR   __reg
		PCICR  __reg
		PCMSK2 __reg
		PCMSK1 __reg
		PCMSK0 __reg
		PCIFR  __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
		PCMSK2: 0x6d, // Pin Change Mask Register 2
		PCMSK1: 0x6c, // Pin Change Mask Register 1
		PCMSK0: 0x6b, // Pin Change Mask Register 0
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPDR __reg
		SPSR __reg
		SPCR __reg
	}{
		SPDR: 0x4e, // SPI Data Register
		SPSR: 0x4d, // SPI Status Register
		SPCR: 0x4c, // SPI Control Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x41, // EEPROM Address Register Bytes
		EEARH: 0x41, // EEPROM Address Register Bytes
		EEDR:  0x40, // EEPROM Data Register
		EECR:  0x3f, // EEPROM Control Register
	}

	// CPU Registers
	CPU = struct {
		PRR    __reg
		OSCCAL __reg
		CLKPR  __reg
		SREG   __reg
		SPL    __reg
		SPH    __reg
		SPMCSR __reg
		MCUCR  __reg
		MCUSR  __reg
		SMCR   __reg
		GPIOR2 __reg
		GPIOR1 __reg
		GPIOR0 __reg
	}{
		PRR:    0x64, // Power Reduction Register
		OSCCAL: 0x66, // Oscillator Calibration Value
		CLKPR:  0x61, // Clock Prescale Register
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		SPMCSR: 0x57, // Store Program Memory Control and Status Register
		MCUCR:  0x55, // MCU Control Register
		MCUSR:  0x54, // MCU Status Register
		SMCR:   0x53, // Sleep Mode Control Register
		GPIOR2: 0x4b, // General Purpose I/O Register 2
		GPIOR1: 0x4a, // General Purpose I/O Register 1
		GPIOR0: 0x3e, // General Purpose I/O Register 0
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL = 0x7 // Brown-out Detector trigger level

	// HIGH
	HIGH_RSTDISBL = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_DWEN     = 0x40 // Debug Wire enable
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON    = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ   = 0x6  // Select boot size
	HIGH_BOOTRST  = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTB0
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for USART: USART
const (
	// UDR0: USART I/O Data Register
	UDR0_UDR0 = 0xff // USART I/O Data bits

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
	UCSR0C_UMSEL0 = 0xc0 // USART Mode Select
	UCSR0C_UPM0   = 0x30 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ0  = 0x6  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity

	// UBRR0L: USART Baud Rate Register Bytes

	// UBRR0H: USART Baud Rate Register Bytes
	UBRR0_UBRR0 = 0xfff // USART Baud rate bits
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM = 0xfe

	// TWBR: TWI Bit Rate register
	TWBR_TWBR = 0xff // TWI Bit rate bits

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

	// TWDR: TWI Data register
	TWDR_TWDR = 0xff // TWI Data bits

	// TWAR: TWI (Slave) Address register
	TWAR_TWA   = 0xfe // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output CompareB Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output CompareA Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80
	TCCR1C_FOC1B = 0x40

	// TCNT1L: Timer/Counter1 Bytes

	// TCNT1H: Timer/Counter1 Bytes
	TCNT1_TCNT1 = 0xffff // Timer/Counter1 bits

	// OCR1AL: Timer/Counter1 Output Compare Register Bytes

	// OCR1AH: Timer/Counter1 Output Compare Register Bytes
	OCR1A_OCR1A = 0xffff // Timer/Counter1 Output Compare A bits

	// OCR1BL: Timer/Counter1 Output Compare Register Bytes

	// OCR1BH: Timer/Counter1 Output Compare Register Bytes
	OCR1B_OCR1B = 0xffff // Timer/Counter1 Output Compare B bits

	// ICR1L: Timer/Counter1 Input Capture Register Bytes

	// ICR1H: Timer/Counter1 Input Capture Register Bytes
	ICR1_ICR1 = 0xffff // Timer/Counter1 Input Capture bits
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

	// TCNT2: Timer/Counter2
	TCNT2_TCNT2 = 0xff // Timer/Counter2 bits

	// OCR2B: Timer/Counter2 Output Compare Register B
	OCR2B_OCR2B = 0xff // Timer/Counter2 Output Compare B bits

	// OCR2A: Timer/Counter2 Output Compare Register A
	OCR2A_OCR2A = 0xff // Timer/Counter2 Output Compare A bits

	// ASSR: Asynchronous Status Register
	ASSR_EXCLK   = 0x40 // Enable External Clock Input
	ASSR_AS2     = 0x20 // Asynchronous Timer/Counter2
	ASSR_TCN2UB  = 0x10 // Timer/Counter2 Update Busy
	ASSR_OCR2AUB = 0x8  // Output Compare Register2 Update Busy
	ASSR_OCR2BUB = 0x4  // Output Compare Register 2 Update Busy
	ASSR_TCR2AUB = 0x2  // Timer/Counter Control Register2 Update Busy
	ASSR_TCR2BUB = 0x1  // Timer/Counter Control Register2 Update Busy
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX   = 0xf  // Analog Channel Selection Bits

	// ADCL: ADC Data Register Bytes

	// ADCH: ADC Data Register Bytes
	ADC_ADC = 0x3ff // ADC Data bits

	// ADCSRA: The ADC Control and Status register A
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC  Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS  = 0x7  // ADC  Prescaler Select Bits

	// ADCSRB: The ADC Control and Status register B
	ADCSRB_ACME = 0x40
	ADCSRB_ADTS = 0x7 // ADC Auto Trigger Source bits

	// DIDR0: Digital Input Disable Register
	DIDR0_ADC5D = 0x20
	DIDR0_ADC4D = 0x10
	DIDR0_ADC3D = 0x8
	DIDR0_ADC2D = 0x4
	DIDR0_ADC1D = 0x2
	DIDR0_ADC0D = 0x1
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

	// DIDR1: Digital Input Disable Register 1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// OCR0B: Timer/Counter0 Output Compare Register
	OCR0B_OCR0B = 0xff // Timer/Counter0 Output Compare B bits

	// OCR0A: Timer/Counter0 Output Compare Register
	OCR0A_OCR0A = 0xff // Timer/Counter0 Output Compare A bits

	// TCNT0: Timer/Counter0
	TCNT0_TCNT0 = 0xff // Timer/Counter0 bits

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

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC1 = 0xc // External Interrupt Sense Control 1 Bits
	EICRA_ISC0 = 0x3 // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT = 0x3 // External Interrupt Request 1 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF = 0x3 // External Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE = 0x7 // Pin Change Interrupt Enables

	// PCMSK2: Pin Change Mask Register 2
	PCMSK2_PCINT = 0xff // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT = 0x7f // Pin Change Enable Masks

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT = 0xff // Pin Change Enable Masks

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF = 0x7 // Pin Change Interrupt Flags
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPDR: SPI Data Register
	SPDR_SPDR = 0xff // SPI Data bits

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

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for EEPROM: EEPROM
const (
	// EEARL: EEPROM Address Register Bytes

	// EEARH: EEPROM Address Register Bytes
	EEAR_EEAR = 0x3ff // EEPROM Address Bits

	// EEDR: EEPROM Data Register
	EEDR_EEDR = 0xff // EEPROM Data Bits

	// EECR: EEPROM Control Register
	EECR_EEPM  = 0x30 // EEPROM Programming Mode Bits
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE = 0x4  // EEPROM Master Write Enable
	EECR_EEPE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTWI    = 0x80 // Power Reduction TWI
	PRR_PRTIM2   = 0x40 // Power Reduction Timer/Counter2
	PRR_PRTIM0   = 0x20 // Power Reduction Timer/Counter0
	PRR_PRTIM1   = 0x8  // Power Reduction Timer/Counter1
	PRR_PRSPI    = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRUSART0 = 0x2  // Power Reduction USART
	PRR_PRADC    = 0x1  // Power Reduction ADC

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS  = 0xf  // Clock Prescaler Select Bits

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// SPMCSR: Store Program Memory Control and Status Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read-While-Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory

	// MCUCR: MCU Control Register
	MCUCR_PUD   = 0x10
	MCUCR_IVSEL = 0x2
	MCUCR_IVCE  = 0x1

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select Bits
	SMCR_SE = 0x1 // Sleep Enable

	// GPIOR2: General Purpose I/O Register 2
	GPIOR2_GPIOR2 = 0xff // General Purpose I/O bits

	// GPIOR1: General Purpose I/O Register 1
	GPIOR1_GPIOR1 = 0xff // General Purpose I/O bits

	// GPIOR0: General Purpose I/O Register 0
	GPIOR0_GPIOR0 = 0xff // General Purpose I/O bits
)
