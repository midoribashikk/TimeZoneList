import java.util.Collections;
import java.util.List;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.TimeZone;
import java.util.Locale;

public class TimeZoneList {
	static List<String> list;

	public static void main(String[] args) {
		TimeZoneList m = new TimeZoneList();
		list = new ArrayList<String>();

		for (String tzid : TimeZone.getAvailableIDs()) {
			m.print(tzid);
		}

		Collections.sort(list);

		for (String elm : list) {
			System.out.println(elm);
		}
	}

	private void print(String tzid) {
		TimeZone tz = TimeZone.getTimeZone(tzid);
		Calendar cal = Calendar.getInstance(tz);

		list.add(tzid + "," + tz.getDisplayName(false, TimeZone.LONG, Locale.US) + "," + minutes(cal, false));
		if (cal.get(Calendar.DST_OFFSET)!=0) {
			list.add(tzid + "," + tz.getDisplayName(true, TimeZone.LONG, Locale.US) + "," + minutes(cal, true));
		}
	}

	private String minutes(Calendar cal, boolean dst) {
		String retstr;
		long offset = cal.get(Calendar.ZONE_OFFSET);

		if (dst) {
			offset += cal.get(Calendar.DST_OFFSET);
		}
		offset /= 1000;

		if (offset<0) {
			retstr = "-";
			offset *= -1;
		} else {
			retstr = "+";
		}

		offset = (offset+30)/60;  // rounding
		retstr += String.valueOf(offset/60) + ":" + String.format("%02d", offset%60);

		return retstr;
	}
}
