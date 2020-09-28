package gst

// #include "gst.go.h"
import "C"

const (
	// ClockFormat is the string used when formatting clock strings
	ClockFormat string = "u:%02u:%02u.%09u"
	// ClockTimeNone means infinite timeout (unsigned representation of -1)
	ClockTimeNone uint64 = 18446744073709551615
)

// BufferingMode is a representation of GstBufferingMode
type BufferingMode int

// Type casts of buffering modes
const (
	BufferingStream    BufferingMode = C.GST_BUFFERING_STREAM
	BufferingDownload  BufferingMode = C.GST_BUFFERING_DOWNLOAD
	BufferingTimeshift BufferingMode = C.GST_BUFFERING_TIMESHIFT
	BufferingLive      BufferingMode = C.GST_BUFFERING_LIVE
)

// Format is a representation of GstFormat.
type Format int

// Type casts of formats
const (
	FormatUndefined Format = C.GST_FORMAT_UNDEFINED
	FormatDefault   Format = C.GST_FORMAT_DEFAULT
	FormatBytes     Format = C.GST_FORMAT_BYTES
	FormatTime      Format = C.GST_FORMAT_TIME
	FormatBuffer    Format = C.GST_FORMAT_BUFFERS
	FormatPercent   Format = C.GST_FORMAT_PERCENT
)

// MessageType is an alias to the C equivalent of GstMessageType.
// See the official documentation for definitions of the messages:
// https://gstreamer.freedesktop.org/documentation/gstreamer/gstmessage.html?gi-language=c#GstMessageType
type MessageType int

// Type casting of GstMessageTypes
// See the official documentation for definitions of the messages:
// https://gstreamer.freedesktop.org/documentation/gstreamer/gstmessage.html?gi-language=c#GstMessageType
const (
	MessageUnknown          MessageType = C.GST_MESSAGE_UNKNOWN
	MessageEOS              MessageType = C.GST_MESSAGE_EOS
	MessageError            MessageType = C.GST_MESSAGE_ERROR
	MessageWarning          MessageType = C.GST_MESSAGE_WARNING
	MessageInfo             MessageType = C.GST_MESSAGE_INFO
	MessageTag              MessageType = C.GST_MESSAGE_TAG
	MessageBuffering        MessageType = C.GST_MESSAGE_BUFFERING
	MessageStateChanged     MessageType = C.GST_MESSAGE_STATE_CHANGED
	MessageStateDirty       MessageType = C.GST_MESSAGE_STATE_DIRTY
	MessageStepDone         MessageType = C.GST_MESSAGE_STEP_DONE
	MessageClockProvide     MessageType = C.GST_MESSAGE_CLOCK_PROVIDE
	MessageClockLost        MessageType = C.GST_MESSAGE_CLOCK_LOST
	MessageNewClock         MessageType = C.GST_MESSAGE_NEW_CLOCK
	MessageStructureChange  MessageType = C.GST_MESSAGE_STRUCTURE_CHANGE
	MessageStreamStatus     MessageType = C.GST_MESSAGE_STREAM_STATUS
	MessageApplication      MessageType = C.GST_MESSAGE_APPLICATION
	MessageElement          MessageType = C.GST_MESSAGE_ELEMENT
	MessageSegmentStart     MessageType = C.GST_MESSAGE_SEGMENT_START
	MessageSegmentDone      MessageType = C.GST_MESSAGE_SEGMENT_DONE
	MessageDurationChanged  MessageType = C.GST_MESSAGE_DURATION_CHANGED
	MessageLatency          MessageType = C.GST_MESSAGE_LATENCY
	MessageAsyncStart       MessageType = C.GST_MESSAGE_ASYNC_START
	MessageAsyncDone        MessageType = C.GST_MESSAGE_ASYNC_DONE
	MessageRequestState     MessageType = C.GST_MESSAGE_REQUEST_STATE
	MessageStepStart        MessageType = C.GST_MESSAGE_STEP_START
	MessageQOS              MessageType = C.GST_MESSAGE_QOS
	MessageProgress         MessageType = C.GST_MESSAGE_PROGRESS
	MessageTOC              MessageType = C.GST_MESSAGE_TOC
	MessageResetTime        MessageType = C.GST_MESSAGE_RESET_TIME
	MessageStreamStart      MessageType = C.GST_MESSAGE_STREAM_START
	MessageNeedContext      MessageType = C.GST_MESSAGE_NEED_CONTEXT
	MessageHaveContext      MessageType = C.GST_MESSAGE_HAVE_CONTEXT
	MessageExtended         MessageType = C.GST_MESSAGE_EXTENDED
	MessageDeviceAdded      MessageType = C.GST_MESSAGE_DEVICE_ADDED
	MessageDeviceRemoved    MessageType = C.GST_MESSAGE_DEVICE_REMOVED
	MessagePropertyNotify   MessageType = C.GST_MESSAGE_PROPERTY_NOTIFY
	MessageStreamCollection MessageType = C.GST_MESSAGE_STREAM_COLLECTION
	MessageStreamsSelected  MessageType = C.GST_MESSAGE_STREAMS_SELECTED
	MessageRedirect         MessageType = C.GST_MESSAGE_REDIRECT
	MessageDeviceChanged    MessageType = C.GST_MESSAGE_DEVICE_CHANGED
	MessageAny              MessageType = C.GST_MESSAGE_ANY
)

// PadDirection is a cast of GstPadDirection to a go type.
type PadDirection int

// Type casting of pad directions
const (
	PadUnknown PadDirection = C.GST_PAD_UNKNOWN // (0) - the direction is unknown
	PadSource  PadDirection = C.GST_PAD_SRC     // (1) - the pad is a source pad
	PadSink    PadDirection = C.GST_PAD_SINK    // (2) - the pad is a sink pad
)

// String implements a Stringer on PadDirection.
func (p PadDirection) String() string {
	switch p {
	case PadUnknown:
		return "Unknown"
	case PadSource:
		return "Src"
	case PadSink:
		return "Sink"
	}
	return ""
}

// PadLinkReturn os a representation of GstPadLinkReturn.
type PadLinkReturn int

// Type casts for PadLinkReturns.
const (
	PadLinkOK             PadLinkReturn = C.GST_PAD_LINK_OK
	PadLinkWrongHierarchy PadLinkReturn = C.GST_PAD_LINK_WRONG_HIERARCHY
	PadLinkWasLinked      PadLinkReturn = C.GST_PAD_LINK_WAS_LINKED
	PadLinkWrongDirection PadLinkReturn = C.GST_PAD_LINK_WRONG_DIRECTION
	PadLinkNoFormat       PadLinkReturn = C.GST_PAD_LINK_NOFORMAT
	PadLinkNoSched        PadLinkReturn = C.GST_PAD_LINK_NOSCHED
	PadLinkRefused        PadLinkReturn = C.GST_PAD_LINK_REFUSED
)

// PadPresence is a cast of GstPadPresence to a go type.
type PadPresence int

// Type casting of pad presences
const (
	PadAlways    PadPresence = C.GST_PAD_ALWAYS    // (0) - the pad is always available
	PadSometimes PadPresence = C.GST_PAD_SOMETIMES // (1) - the pad will become available depending on the media stream
	PadRequest   PadPresence = C.GST_PAD_REQUEST   // (2) - the pad is only available on request with gst_element_request_pad.
)

// String implements a stringer on PadPresence.
func (p PadPresence) String() string {
	switch p {
	case PadAlways:
		return "Always"
	case PadSometimes:
		return "Sometimes"
	case PadRequest:
		return "Request"
	}
	return ""
}

// State is a type cast of the C GstState
type State int

// Type casting for GstStates
const (
	VoidPending  State = C.GST_STATE_VOID_PENDING // (0) – no pending state.
	StateNull    State = C.GST_STATE_NULL         // (1) – the NULL state or initial state of an element.
	StateReady   State = C.GST_STATE_READY        // (2) – the element is ready to go to PAUSED.
	StatePaused  State = C.GST_STATE_PAUSED       // (3) – the element is PAUSED, it is ready to accept and process data. Sink elements however only accept one buffer and then block.
	StatePlaying State = C.GST_STATE_PLAYING      // (4) – the element is PLAYING, the GstClock is running and the data is flowing.
)

// String returns the string representation of this state.
func (s State) String() string {
	return C.GoString(C.gst_element_state_get_name((C.GstState)(s)))
}

// SeekFlags is a representation of GstSeekFlags.
type SeekFlags int

// Type casts of SeekFlags
const (
	SeekFlagNone        SeekFlags = C.GST_SEEK_FLAG_NONE
	SeekFlagFlush       SeekFlags = C.GST_SEEK_FLAG_FLUSH
	SeekFlagAccurate    SeekFlags = C.GST_SEEK_FLAG_ACCURATE
	SeekFlagKeyUnit     SeekFlags = C.GST_SEEK_FLAG_KEY_UNIT
	SeekFlagSegment     SeekFlags = C.GST_SEEK_FLAG_SEGMENT
	SeekFlagSkip        SeekFlags = C.GST_SEEK_FLAG_SKIP
	SeekFlagSnapBefore  SeekFlags = C.GST_SEEK_FLAG_SNAP_BEFORE
	SeekFlagSnapAfter   SeekFlags = C.GST_SEEK_FLAG_SNAP_AFTER
	SeekFlagSnapNearest SeekFlags = C.GST_SEEK_FLAG_SNAP_NEAREST
)

// SeekType is a representation of GstSeekType.
type SeekType int

// Type casts of seek types
const (
	SeekTypeNone SeekType = C.GST_SEEK_TYPE_NONE
	SeekTypeSet  SeekType = C.GST_SEEK_TYPE_SET
	SeekTypeEnd  SeekType = C.GST_SEEK_TYPE_END
)

// StateChangeReturn is a representation of GstStateChangeReturn.
type StateChangeReturn int

// Type casts of state change returns
const (
	StateChangeFailure   StateChangeReturn = C.GST_STATE_CHANGE_FAILURE
	StateChangeSuccess   StateChangeReturn = C.GST_STATE_CHANGE_SUCCESS
	StateChangeAsync     StateChangeReturn = C.GST_STATE_CHANGE_ASYNC
	StateChangeNoPreroll StateChangeReturn = C.GST_STATE_CHANGE_NO_PREROLL
)

// ElementFlags casts C GstElementFlags to a go type
type ElementFlags int

// Type casting of element flags
const (
	ElementFlagLockedState  ElementFlags = C.GST_ELEMENT_FLAG_LOCKED_STATE  // (16) – ignore state changes from parent
	ElementFlagSink         ElementFlags = C.GST_ELEMENT_FLAG_SINK          // (32) – the element is a sink
	ElementFlagSource       ElementFlags = C.GST_ELEMENT_FLAG_SOURCE        // (64) – the element is a source.
	ElementFlagProvideClock ElementFlags = C.GST_ELEMENT_FLAG_PROVIDE_CLOCK // (128) – the element can provide a clock
	ElementFlagRequireClock ElementFlags = C.GST_ELEMENT_FLAG_REQUIRE_CLOCK // (256) – the element requires a clock
	ElementFlagIndexable    ElementFlags = C.GST_ELEMENT_FLAG_INDEXABLE     // (512) – the element can use an index
	ElementFlagLast         ElementFlags = C.GST_ELEMENT_FLAG_LAST          // (16384) – offset to define more flags
)

// MiniObjectFlags casts GstMiniObjectFlags to a go type.
type MiniObjectFlags int

// Type casting of mini-object flags
const (
	MiniObjectFlagLockable     MiniObjectFlags = C.GST_MINI_OBJECT_FLAG_LOCKABLE      // (1) – the object can be locked and unlocked with gst_mini_object_lock and gst_mini_object_unlock.
	MiniObjectFlagLockReadOnly MiniObjectFlags = C.GST_MINI_OBJECT_FLAG_LOCK_READONLY // (2) – the object is permanently locked in READONLY mode. Only read locks can be performed on the object.
	MiniObjectFlagMayBeLeaked  MiniObjectFlags = C.GST_MINI_OBJECT_FLAG_MAY_BE_LEAKED // (4) – the object is expected to stay alive even after gst_deinit has been called and so should be ignored by leak detection tools. (Since: 1.10)
	MiniObjectFlagLast         MiniObjectFlags = C.GST_MINI_OBJECT_FLAG_LAST          // (16) – first flag that can be used by subclasses.
)

// FlowReturn is go type casting for GstFlowReturn.
type FlowReturn int

// Type casting of the GstFlowReturn types. Custom ones are omitted for now.
const (
	FlowOK            FlowReturn = C.GST_FLOW_OK             // Data passing was ok
	FlowNotLinked     FlowReturn = C.GST_FLOW_NOT_LINKED     // Pad is not linked
	FlowFlushing      FlowReturn = C.GST_FLOW_FLUSHING       // Pad is flushing
	FlowEOS           FlowReturn = C.GST_FLOW_EOS            // Pad is EOS
	FlowNotNegotiated FlowReturn = C.GST_FLOW_NOT_NEGOTIATED // Pad is not negotiated
	FlowError         FlowReturn = C.GST_FLOW_ERROR          // Some (fatal) error occurred
	FlowNotSupported  FlowReturn = C.GST_FLOW_NOT_SUPPORTED  // The operation is not supported.
)

// MapFlags is a go casting of GstMapFlags
type MapFlags int

// Type casting of the map flag types
const (
	MapRead     MapFlags = C.GST_MAP_READ      //  (1) – map for read access
	MapWrite    MapFlags = C.GST_MAP_WRITE     // (2) - map for write access
	MapFlagLast MapFlags = C.GST_MAP_FLAG_LAST // (65536) – first flag that can be used for custom purposes
)

// StreamStatusType represents a type of change in a stream's status
type StreamStatusType int

// Type castings of the stream status types
const (
	StreamStatusCreate  = C.GST_STREAM_STATUS_TYPE_CREATE  // (0) – A new thread need to be created.
	StreamStatusEnter   = C.GST_STREAM_STATUS_TYPE_ENTER   // (1) – a thread entered its loop function
	StreamStatusLeave   = C.GST_STREAM_STATUS_TYPE_LEAVE   // (2) – a thread left its loop function
	StreamStatusDestroy = C.GST_STREAM_STATUS_TYPE_DESTROY // (3) – a thread is destroyed
	StreamStatusStart   = C.GST_STREAM_STATUS_TYPE_START   // (8) – a thread is started
	StreamStatusPause   = C.GST_STREAM_STATUS_TYPE_PAUSE   // (9) – a thread is paused
	StreamStatusStop    = C.GST_STREAM_STATUS_TYPE_STOP    // (10) – a thread is stopped
)

func (s StreamStatusType) String() string {
	switch s {
	case StreamStatusCreate:
		return "A new thread needs to be created"
	case StreamStatusEnter:
		return "A thread has entered its loop function"
	case StreamStatusLeave:
		return "A thread has left its loop function"
	case StreamStatusDestroy:
		return "A thread has been destroyed"
	case StreamStatusStart:
		return "A thread has started"
	case StreamStatusPause:
		return "A thread has paused"
	case StreamStatusStop:
		return "A thread has stopped"
	}
	return ""
}