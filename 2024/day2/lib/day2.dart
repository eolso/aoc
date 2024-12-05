import 'dart:io';

List<List<int>> readReports(String file) {
  List<String> lines = File("input.txt").readAsLinesSync();

  List<List<int>> reports = [];
  for (var line in lines) {
    var splitLine = line.split(" ");

    List<int> report = [];
    for (var a in splitLine) {
      report.add(int.parse(a));
    }

    reports.add(report);
  }

  return reports;
}

bool isSafe(List<int> values, int maxStep) {
  return (isIncreasing(values) || isDecreasing(values)) &&
      isStepped(values, maxStep);
}

bool isSafeish(List<int> values, int maxStep) {
  if (isSafe(values, maxStep)) {
    return true;
  }

  for (var i = 0; i < values.length; i++) {
    var subValues = List<int>.from(values);
    subValues.removeAt(i);

    if (isSafe(subValues, maxStep)) {
      return true;
    }
  }

  return false;
}

bool isIncreasing(List<int> values) {
  if (values.length <= 1) {
    return true;
  }

  var lastValue = values[0];
  for (var i = 1; i < values.length; i++) {
    if (values[i] <= lastValue) {
      return false;
    }

    lastValue = values[i];
  }

  return true;
}

bool isDecreasing(List<int> values) {
  if (values.length <= 1) {
    return true;
  }

  var lastValue = values[0];
  for (var i = 1; i < values.length; i++) {
    if (values[i] >= lastValue) {
      return false;
    }

    lastValue = values[i];
  }

  return true;
}

bool isStepped(List<int> values, int maxStep) {
  if (values.length <= 1) {
    return true;
  }

  var lastValue = values[0];
  for (var i = 1; i < values.length; i++) {
    if (values[i] == lastValue || (values[i] - lastValue).abs() > maxStep) {
      return false;
    }

    lastValue = values[i];
  }

  return true;
}
