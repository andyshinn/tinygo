// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega32U2.atdf, see http://packs.download.atmel.com/

// +build avr,atmega32u2

// Device information for the ATmega32U2.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega32U2"
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
	IRQ_PCINT1       = 10 // Pin Change Interrupt Request 1
	IRQ_USB_GEN      = 11 // USB General Interrupt Request
	IRQ_USB_COM      = 12 // USB Endpoint/Pipe Interrupt Communication Request
	IRQ_WDT          = 13 // Watchdog Time-out Interrupt
	IRQ_TIMER1_CAPT  = 14 // Timer/Counter2 Capture Event
	IRQ_TIMER1_COMPA = 15 // Timer/Counter2 Compare Match B
	IRQ_TIMER1_COMPB = 16 // Timer/Counter2 Compare Match B
	IRQ_TIMER1_COMPC = 17 // Timer/Counter2 Compare Match C
	IRQ_TIMER1_OVF   = 18 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 19 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 20 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_OVF   = 21 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 22 // SPI Serial Transfer Complete
	IRQ_USART1_RX    = 23 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 24 // USART1 Data register Empty
	IRQ_USART1_TX    = 25 // USART1, Tx Complete
	IRQ_ANALOG_COMP  = 26 // Analog Comparator
	IRQ_EE_READY     = 27 // EEPROM Ready
	IRQ_SPM_READY    = 28 // Store Program Memory Read
	IRQ_max          = 28 // Highest interrupt number on this device.
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

	// I/O Port
	PORT = struct {
		PORTB __reg
		DDRB  __reg
		PINB  __reg
		PORTD __reg
		DDRD  __reg
		PIND  __reg
		PORTC __reg
		DDRC  __reg
		PINC  __reg
	}{
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
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
		GTCCR  __reg
	}{
		OCR0B:  0x48, // Timer/Counter0 Output Compare Register
		OCR0A:  0x47, // Timer/Counter0 Output Compare Register
		TCNT0:  0x46, // Timer/Counter0
		TCCR0B: 0x45, // Timer/Counter Control Register B
		TCCR0A: 0x44, // Timer/Counter  Control Register A
		TIMSK0: 0x6e, // Timer/Counter0 Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter0 Interrupt Flag register
		GTCCR:  0x43, // General Timer/Counter Control Register
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
		TIMSK1 __reg
		TIFR1  __reg
	}{
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

	// Phase Locked Loop
	PLL = struct {
		PLLCSR __reg
	}{
		PLLCSR: 0x49, // PLL Status and Control register
	}

	// USB Device Registers
	USB_DEVICE = struct {
		UPOE    __reg
		UEINT   __reg
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
		USBCON  __reg
		REGCR   __reg
	}{
		UPOE:    0xfb, // USB Software Output Enable register
		UEINT:   0xf4, // USB Endpoint Number Interrupt Register
		UEBCLX:  0xf2, // USB Endpoint Byte Count Register
		UEDATX:  0xf1, // USB Data Endpoint
		UEIENX:  0xf0, // USB Endpoint Interrupt Enable Register
		UESTA1X: 0xef, // USB Endpoint Status 1 Register
		UESTA0X: 0xee, // USB Endpoint Status 0 Register
		UECFG1X: 0xed, // USB Endpoint Configuration 1 Register
		UECFG0X: 0xec, // USB Endpoint Configuration 0 Register
		UECONX:  0xeb, // USB Endpoint Control Register
		UERST:   0xea, // USB Endpoint Reset Register
		UENUM:   0xe9, // USB Endpoint Number
		UEINTX:  0xe8, // USB Endpoint Interrupt Register
		UDMFN:   0xe6, // USB Device Micro Frame Number
		UDFNUML: 0xe4, // USB Device Frame Number High Register
		UDFNUMH: 0xe4, // USB Device Frame Number High Register
		UDADDR:  0xe3, // USB Device Address Register
		UDIEN:   0xe2, // USB Device Interrupt Enable Register
		UDINT:   0xe1, // USB Device Interrupt Register
		UDCON:   0xe0, // USB Device Control Registers
		USBCON:  0xd8, // USB General Control Register
		REGCR:   0x63, // Regulator Control Register
	}

	// CPU Registers
	CPU = struct {
		SREG    __reg
		SPL     __reg
		SPH     __reg
		MCUCR   __reg
		MCUSR   __reg
		OSCCAL  __reg
		CLKPR   __reg
		SMCR    __reg
		EIND    __reg
		GPIOR2  __reg
		GPIOR1  __reg
		GPIOR0  __reg
		PRR1    __reg
		PRR0    __reg
		CLKSTA  __reg
		CLKSEL1 __reg
		CLKSEL0 __reg
		DWDR    __reg
	}{
		SREG:    0x5f, // Status Register
		SPL:     0x5d, // Stack Pointer
		SPH:     0x5d, // Stack Pointer
		MCUCR:   0x55, // MCU Control Register
		MCUSR:   0x54, // MCU Status Register
		OSCCAL:  0x66, // Oscillator Calibration Value
		CLKPR:   0x61,
		SMCR:    0x53, // Sleep Mode Control Register
		EIND:    0x5c, // Extended Indirect Register
		GPIOR2:  0x4b, // General Purpose IO Register 2
		GPIOR1:  0x4a, // General Purpose IO Register 1
		GPIOR0:  0x3e, // General Purpose IO Register 0
		PRR1:    0x65, // Power Reduction Register1
		PRR0:    0x64, // Power Reduction Register0
		CLKSTA:  0xd2,
		CLKSEL1: 0xd1,
		CLKSEL0: 0xd0,
		DWDR:    0x51, // debugWire communication register
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EICRB  __reg
		EIMSK  __reg
		EIFR   __reg
		PCMSK0 __reg
		PCMSK1 __reg
		PCIFR  __reg
		PCICR  __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register A
		EICRB:  0x6a, // External Interrupt Control Register B
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCMSK0: 0x6b, // Pin Change Mask Register 0
		PCMSK1: 0x6c, // Pin Change Mask Register 1
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
	}

	// USART
	USART = struct {
		UDR1   __reg
		UCSR1A __reg
		UCSR1B __reg
		UCSR1C __reg
		UCSR1D __reg
		UBRR1L __reg
		UBRR1H __reg
	}{
		UDR1:   0xce, // USART I/O Data Register
		UCSR1A: 0xc8, // USART Control and Status Register A
		UCSR1B: 0xc9, // USART Control and Status Register B
		UCSR1C: 0xca, // USART Control and Status Register C
		UCSR1D: 0xcb, // USART Control and Status Register D
		UBRR1L: 0xcc, // USART Baud Rate Register  Bytes
		UBRR1H: 0xcc, // USART Baud Rate Register  Bytes
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
		WDTCKD __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
		WDTCKD: 0x62, // Watchdog Timer Clock Divider
	}

	// Analog Comparator
	AC = struct {
		ACSR  __reg
		ACMUX __reg
		DIDR1 __reg
	}{
		ACSR:  0x50, // Analog Comparator Control And Status Register
		ACMUX: 0x7d, // Analog Comparator Input Multiplexer
		DIDR1: 0x7f,
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL = 0x7 // Brown-out Detector trigger level
	EXTENDED_HWBE     = 0x8 // Hardware Boot Enable

	// HIGH
	HIGH_DWEN     = 0x80 // Debug Wire enable
	HIGH_RSTDISBL = 0x40 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON    = 0x10 // Watchdog timer always on
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ   = 0x6  // Select Boot Size
	HIGH_BOOTRST  = 0x1  // Boot Reset vector Enabled

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

// Bitfields for PORT: I/O Port
const (
	// PORTC: Port C Data Register
	PORTC_PORTC = 0xf0 // Port C Data Register bits
	PORTC_PORTC = 0x7  // Port C Data Register bits

	// DDRC: Port C Data Direction Register
	DDRC_DDC = 0xf0 // Port C Data Direction Register bits
	DDRC_DDC = 0x7  // Port C Data Direction Register bits

	// PINC: Port C Input Pins
	PINC_PINC = 0xf0 // Port C Input Pins bits
	PINC_PINC = 0x7  // Port C Input Pins bits
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

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM     = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_PSRSYNC = 0x1  // Prescaler Reset Timer/Counter1 and Timer/Counter0
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
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

// Bitfields for PLL: Phase Locked Loop
const (
	// PLLCSR: PLL Status and Control register
	PLLCSR_PLLP  = 0x1c // PLL prescaler Bits
	PLLCSR_PLLE  = 0x2  // PLL Enable Bit
	PLLCSR_PLOCK = 0x1  // PLL Lock Status Bit
)

// Bitfields for USB_DEVICE: USB Device Registers
const (
	// UPOE: USB Software Output Enable register
	UPOE_UPWE  = 0xc0 // USB Buffers Direct Drive enable configuration
	UPOE_UPDRV = 0x30 // USB direct drive values
	UPOE_DPI   = 0x2  // D+ Input value
	UPOE_DMI   = 0x1  // D- Input value

	// UEINT: USB Endpoint Number Interrupt Register
	UEINT_EPINT = 0x1f // Byte Count bits

	// UEBCLX: USB Endpoint Byte Count Register
	UEBCLX_BYCT = 0xff // Byte Count bits

	// UEDATX: USB Data Endpoint
	UEDATX_DAT = 0xff // Data bits

	// UEIENX: USB Endpoint Interrupt Enable Register
	UEIENX_FLERRE   = 0x80 // Flow Error Interrupt Enable Flag
	UEIENX_NAKINE   = 0x40 // NAK IN Interrupt Enable Bit
	UEIENX_NAKOUTE  = 0x10 // NAK OUT Interrupt Enable Bit
	UEIENX_RXSTPE   = 0x8  // Received SETUP Interrupt Enable Flag
	UEIENX_RXOUTE   = 0x4  // Received OUT Data Interrupt Enable Flag
	UEIENX_STALLEDE = 0x2  // Stalled Interrupt Enable Flag
	UEIENX_TXINE    = 0x1  // Transmitter Ready Interrupt Enable Flag

	// UESTA1X: USB Endpoint Status 1 Register
	UESTA1X_CTRLDIR = 0x4 // Control Direction
	UESTA1X_CURRBK  = 0x3 // Current Bank

	// UESTA0X: USB Endpoint Status 0 Register
	UESTA0X_CFGOK   = 0x80 // Configuration Status Flag
	UESTA0X_OVERFI  = 0x40 // Overflow Error Interrupt Flag
	UESTA0X_UNDERFI = 0x20 // Underflow Error Interrupt Flag
	UESTA0X_DTSEQ   = 0xc  // Data Toggle Sequencing Flag
	UESTA0X_NBUSYBK = 0x3  // Busy Bank Flag

	// UECFG1X: USB Endpoint Configuration 1 Register
	UECFG1X_EPSIZE = 0x70 // Endpoint Size Bits
	UECFG1X_EPBK   = 0xc  // Endpoint Bank Bits
	UECFG1X_ALLOC  = 0x2  // Endpoint Allocation Bit

	// UECFG0X: USB Endpoint Configuration 0 Register
	UECFG0X_EPTYPE = 0xc0 // Endpoint Type Bits
	UECFG0X_EPDIR  = 0x1  // Endpoint Direction Bit

	// UECONX: USB Endpoint Control Register
	UECONX_STALLRQ  = 0x20 // STALL Request Handshake Bit
	UECONX_STALLRQC = 0x10 // STALL Request Clear Handshake Bit
	UECONX_RSTDT    = 0x8  // Reset Data Toggle Bit
	UECONX_EPEN     = 0x1  // Endpoint Enable Bit

	// UERST: USB Endpoint Reset Register
	UERST_EPRST = 0x1f // Endpoint FIFO Reset Bits

	// UENUM: USB Endpoint Number
	UENUM_EPNUM = 0x7 // Endpoint Number bits

	// UEINTX: USB Endpoint Interrupt Register
	UEINTX_FIFOCON  = 0x80 // FIFO Control Bit
	UEINTX_NAKINI   = 0x40 // NAK IN Received Interrupt Flag
	UEINTX_RWAL     = 0x20 // Read/Write Allowed Flag
	UEINTX_NAKOUTI  = 0x10 // NAK OUT Received Interrupt Flag
	UEINTX_RXSTPI   = 0x8  // Received SETUP Interrupt Flag
	UEINTX_RXOUTI   = 0x4  // Received OUT Data Interrupt Flag
	UEINTX_STALLEDI = 0x2  // STALLEDI Interrupt Flag
	UEINTX_TXINI    = 0x1  // Transmitter Ready Interrupt Flag

	// UDMFN: USB Device Micro Frame Number
	UDMFN_FNCERR = 0x10 // Frame Number CRC Error Flag

	// UDFNUML: USB Device Frame Number High Register

	// UDFNUMH: USB Device Frame Number High Register
	UDFNUM_FNUM = 0x7ff // Frame Number Upper Flag

	// UDADDR: USB Device Address Register
	UDADDR_ADDEN = 0x80 // Address Enable Bit
	UDADDR_UADD  = 0x7f // USB Address Bits

	// UDIEN: USB Device Interrupt Enable Register
	UDIEN_UPRSME  = 0x40 // Upstream Resume Interrupt Enable Bit
	UDIEN_EORSME  = 0x20 // End Of Resume Interrupt Enable Bit
	UDIEN_WAKEUPE = 0x10 // Wake-up CPU Interrupt Enable Bit
	UDIEN_EORSTE  = 0x8  // End Of Reset Interrupt Enable Bit
	UDIEN_SOFE    = 0x4  // Start Of Frame Interrupt Enable Bit
	UDIEN_SUSPE   = 0x1  // Suspend Interrupt Enable Bit

	// UDINT: USB Device Interrupt Register
	UDINT_UPRSMI  = 0x40 // Upstream Resume Interrupt Flag
	UDINT_EORSMI  = 0x20 // End Of Resume Interrupt Flag
	UDINT_WAKEUPI = 0x10 // Wake-up CPU Interrupt Flag
	UDINT_EORSTI  = 0x8  // End Of Reset Interrupt Flag
	UDINT_SOFI    = 0x4  // Start Of Frame Interrupt Flag
	UDINT_SUSPI   = 0x1  // Suspend Interrupt Flag

	// UDCON: USB Device Control Registers
	UDCON_RSTCPU = 0x4 // USB Reset CPU Bit
	UDCON_RMWKUP = 0x2 // Remote Wake-up Bit
	UDCON_DETACH = 0x1 // Detach Bit

	// USBCON: USB General Control Register
	USBCON_USBE   = 0x80 // USB macro Enable Bit
	USBCON_FRZCLK = 0x20 // Freeze USB Clock Bit

	// REGCR: Regulator Control Register
	REGCR_REGDIS = 0x1 // Regulator Disable
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
	MCUCR_PUD   = 0x10 // Pull-up disable
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_USBRF = 0x20 // USB reset flag
	MCUSR_WDRF  = 0x8  // Watchdog Reset Flag
	MCUSR_BORF  = 0x4  // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2  // External Reset Flag
	MCUSR_PORF  = 0x1  // Power-on reset flag

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
	PRR1_PRUSART1 = 0x1  // Power Reduction USART1

	// PRR0: Power Reduction Register0
	PRR0_PRTIM0 = 0x20 // Power Reduction Timer/Counter0
	PRR0_PRTIM1 = 0x8  // Power Reduction Timer/Counter1
	PRR0_PRSPI  = 0x4  // Power Reduction Serial Peripheral Interface

	// CLKSTA
	CLKSTA_RCON  = 0x2
	CLKSTA_EXTON = 0x1

	// CLKSEL1
	CLKSEL1_RCCKSEL = 0xf0
	CLKSEL1_EXCKSEL = 0xf

	// CLKSEL0
	CLKSEL0_RCSUT = 0xc0
	CLKSEL0_EXSUT = 0x30
	CLKSEL0_RCE   = 0x8
	CLKSEL0_EXTE  = 0x4
	CLKSEL0_CLKS  = 0x1
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

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT = 0xff // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT = 0x1f

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF = 0x3 // Pin Change Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE = 0x3 // Pin Change Interrupt Enables
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

	// UCSR1D: USART Control and Status Register D
	UCSR1D_CTSEN = 0x2 // CTS Enable
	UCSR1D_RTSEN = 0x1 // RTS Enable
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable

	// WDTCKD: Watchdog Timer Clock Divider
	WDTCKD_WDEWIF = 0x8 // Watchdog Early Warning Interrupt Flag
	WDTCKD_WDEWIE = 0x4 // Watchdog Early Warning Interrupt Enable
	WDTCKD_WCLKD  = 0x3 // Watchdog Timer Clock Dividers
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

	// ACMUX: Analog Comparator Input Multiplexer
	ACMUX_CMUX = 0x7 // Analog Comparator Selection Bits

	// DIDR1
	DIDR1_AIN7D = 0x80 // AIN7 Digital Input Disable
	DIDR1_AIN6D = 0x40 // AIN6 Digital Input Disable
	DIDR1_AIN5D = 0x20 // AIN5 Digital Input Disable
	DIDR1_AIN4D = 0x10 // AIN4 Digital Input Disable
	DIDR1_AIN3D = 0x8  // AIN3 Digital Input Disable
	DIDR1_AIN2D = 0x4  // AIN2 Digital Input Disable
	DIDR1_AIN1D = 0x2  // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1  // AIN0 Digital Input Disable
)
