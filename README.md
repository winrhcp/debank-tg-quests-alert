# Debank TG Quests Alert

This project monitors the Debank quests and sends notifications to a Telegram channel when new quests are available.

Follow Us
Stay updated and follow our Telegram channel: https://t.me/debank_quests_alert

# Configure Environment Variables
Create a .env file in the root directory with the following content:
API_URL=https://api.debank.com/quest/list?limit=50&status=hot
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHANNEL_ID=@your_channel_id
