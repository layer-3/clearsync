// Generic messages
export const CHANNEL_FINALIZED = 'Channel finalized.';
export const CHANNEL_NOT_FINALIZED = 'Channel not finalized.';
export const CHANNEL_NOT_OPEN = 'Channel not open.';
export const INVALID_SIGNATURES = 'Invalid signatures';
export const NO_ONGOING_CHALLENGE = 'No ongoing challenge.';
export const TURN_NUM_RECORD_DECREASED = 'turnNumRecord decreased.';
export const TURN_NUM_RECORD_NOT_INCREASED = 'turnNumRecord not increased.';
export const UNACCEPTABLE_WHO_SIGNED_WHAT = 'Unacceptable whoSignedWhat array';
export const WHO_SIGNED_WHAT_WRONG_LENGTH = '|whoSignedWhat|!=nParticipants';
export const WRONG_CHANNEL_STORAGE = 'status(ChannelData)!=storage';
export const INVALID_SIGNATURE = 'Invalid signature';

// Function-specific messages
export const CHALLENGER_NON_PARTICIPANT = 'Challenger is not a participant';
export const RESPONSE_UNAUTHORIZED = 'Signer not authorized mover';
export const WRONG_REFUTATION_SIGNATURE = 'Refutation state not signed by challenger';
export const NONFINAL_STATE = 'State must be final';

// Application-specifis messages
export const COUNTING_APP_INVALID_TRANSITION = 'Counter must be incremented';

// Turn Taking
export const INVALID_SIGNED_BY = 'Invalid signedBy';
export const INVALID_NUMBER_OF_PROOF_STATES = 'Invalid number of proof states';
export const TOO_MANY_PARTICIPANTS = 'Too many participants';
export const WRONG_TURN_NUM = 'Wrong variablePart.turnNum';
export const NOT_UNANIMOUS = '!unanimous';
export const PROOF_SUPPLIED = '|proof|!=0';
