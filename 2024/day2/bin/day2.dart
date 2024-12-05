import 'package:day2/day2.dart' as day2;

void main(List<String> arguments) {
  var reports = day2.readReports("input.txt");

  int safeReports = 0;
  int safeishReports = 0;
  for (var report in reports) {
    if (day2.isSafe(report, 3)) {
      safeReports += 1;
    }

    if (day2.isSafeish(report, 3)) {
      safeishReports += 1;
    }
  }

  print("safe reports: $safeReports");
  print("safeish reports: $safeishReports");
}
